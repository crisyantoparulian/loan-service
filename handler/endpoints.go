package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/crisyantoparulian/loansvc/generated"
	"github.com/crisyantoparulian/loansvc/models"
	"github.com/crisyantoparulian/loansvc/repository"
	apperror "github.com/crisyantoparulian/loansvc/utils/app_error"
	"github.com/crisyantoparulian/loansvc/utils/constant"
	httphelper "github.com/crisyantoparulian/loansvc/utils/http_helper"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// API for getting list of loans
// (GET /loans)
func (s *Server) GetLoans(c echo.Context, params generated.GetLoansParams) error {
	ctx := c.Request().Context()

	input := repository.GetLoansInput{}.FromParam(params)

	loans, total, err := s.Repository.GetListLoans(ctx, input)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	resp := mappingToLoanListResponse(loans, total, input)

	return httphelper.HttpSuccessOk(c, "success get list loan", resp)
}

// API for submitting a new loan proposal
// (POST /loans)
func (s *Server) SubmitLoanProposal(c echo.Context, params generated.SubmitLoanProposalParams) error {

	ctx := c.Request().Context()
	payload := generated.SubmitLoanProposalJSONRequestBody{}
	if err := s.bindAndValidatePayload(c, &payload); err != nil {
		return err
	}

	borrower, err := s.Repository.GetBorrowerByIDWithDetail(ctx, params.LoginUserId)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	loanID := uuid.New()
	err = s.Repository.CreateLoan(ctx, &models.Loan{
		ID:              loanID,
		LoanCode:        generateLoanCode(borrower.Type),
		BorrowerID:      borrower.ID,
		PrincipalAmount: payload.PrincipalAmount,
		ROI:             payload.Roi,
		Rate:            payload.Rate,
		Status:          constant.LOAN_STATUS_PROPOSED.String(),
	})
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	resp := generated.SubmitLoanResponseData{
		LoanId: loanID,
	}

	return httphelper.HttpSuccessCreated(c, "success create loan", resp)
}

// API for getting loan details
// (GET /loans/{loanID})
func (s *Server) GetLoanDetails(c echo.Context, loanID openapi_types.UUID, params generated.GetLoanDetailsParams) error {
	ctx := c.Request().Context()

	loan, err := s.Repository.GetLoanByIDWithDetailSQL(ctx, loanID)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// mapping the response
	resp := mappingToDetailLoanResponse(loan)

	return httphelper.HttpSuccessOk(c, "success get detail", resp)
}

// Approve loan proposal by admin
// (PUT /loans/{loanID}/approve)
func (s *Server) ApproveLoan(c echo.Context, loanID openapi_types.UUID, params generated.ApproveLoanParams) error {
	ctx := c.Request().Context()

	loan, err := s.Repository.GetLoanByID(ctx, loanID)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// validate loan status
	if constant.LoanStatus(loan.Status) != constant.LOAN_STATUS_PROPOSED {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(fmt.Errorf("loan status must %s", constant.LOAN_STATUS_PROPOSED), http.StatusUnprocessableEntity))
	}

	visits, err := s.Repository.GetVisits(ctx, repository.GetVisitInput{LoandID: loanID, Limit: 1})
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// validate have minimum 1 visit
	if len(visits) == 0 {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(errors.New("no visit exist for this loan"), http.StatusUnprocessableEntity))
	}

	now := time.Now()
	loan.Status = constant.LOAN_STATUS_APPROVED.String()
	loan.ApprovalDate = &now
	err = s.Repository.ApproveLoan(ctx, loan)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	return httphelper.HttpSuccessOk(c, "success approve", nil)
}

// Create loan disbursement
// (POST /loans/{loanID}/disburse)
func (s *Server) DisburseLoan(c echo.Context, loanID openapi_types.UUID, params generated.DisburseLoanParams) error {
	ctx := c.Request().Context()

	payload := generated.DisburseLoanJSONRequestBody{}
	if err := s.bindAndValidatePayload(c, &payload); err != nil {
		return err
	}

	loan, err := s.Repository.GetLoanByID(ctx, loanID)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// validate loan status
	if constant.LoanStatus(loan.Status) != constant.LOAN_STATUS_INVESTED {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(fmt.Errorf("loan status must %s", constant.LOAN_STATUS_INVESTED), http.StatusUnprocessableEntity))
	}

	// this will validate employee exist on DB
	employee, err := s.Repository.GetEmployeeByID(ctx, payload.FieldOfficerId)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	if employee.Role != constant.ROLE_FIELD_OFFICER.String() {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(errors.New("invalid field_officer_id role"), http.StatusUnprocessableEntity))
	}

	// prepare model disbursement & documents
	disbursement := models.Disbursement{
		ID:                 uuid.New(),
		LoanID:             loanID,
		FieldOfficerID:     payload.FieldOfficerId,
		DateOfDisbursement: payload.DateOfDisbursement,
	}
	documents := []models.Document{
		{
			ID:            uuid.New(),
			ReferenceID:   disbursement.ID,
			ReferenceType: "table_disbursements", // table reference, TODO: update using enum constant
			DocumentType:  "loan_aggrement",      // type docs, TODO: update using enum constant
			FileURL:       payload.LoanAggrementUrl,
		},
	}

	// update loan status
	loan.Status = constant.LOAN_STATUS_DISBURSED.String()
	err = s.Repository.CreateDisbursement(ctx, &disbursement, documents, loan)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	return httphelper.HttpSuccessCreated(c, "success disburse loan", nil)
}

// Create investment on loan
// (POST /loans/{loanID}/invest)
func (s *Server) InvestLoan(c echo.Context, loanID openapi_types.UUID, params generated.InvestLoanParams) error {
	ctx := c.Request().Context()

	payload := generated.InvestLoanJSONRequestBody{}
	if err := s.bindAndValidatePayload(c, &payload); err != nil {
		return err
	}

	loan, err := s.Repository.GetLoanByID(ctx, loanID)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// validate loan status
	if constant.LoanStatus(loan.Status) != constant.LOAN_STATUS_APPROVED {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(fmt.Errorf("loan status must %s", constant.LOAN_STATUS_APPROVED), http.StatusUnprocessableEntity))
	}
	// validate amount total investment
	if payload.Amount > (loan.PrincipalAmount - loan.TotalInvestmentAmount) {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(fmt.Errorf("amount exceed total maximum investment, %f left", loan.TotalInvestmentAmount),
			http.StatusUnprocessableEntity))
	}

	investments, err := s.Repository.GetLoanInvestmentByInvestorID(ctx, loanID, params.LoginUserId)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// validate the investor never invest on the loan
	if len(investments) > 0 {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(errors.New("already invest on this loan"), http.StatusUnprocessableEntity))
	}

	// prepare model investment
	investment := &models.Investment{
		ID:         uuid.New(),
		LoanID:     loanID,
		InvestorID: params.LoginUserId,
		Amount:     payload.Amount,
	}

	// set loan field for update
	if (loan.TotalInvestmentAmount + payload.Amount) == loan.PrincipalAmount {
		loan.Status = constant.LOAN_STATUS_INVESTED.String()
	}
	loan.TotalInvestmentAmount += investment.Amount
	loan.TotalInvestor += 1

	err = s.Repository.CreateInvestment(ctx, investment, loan)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// TODO: pulish event for generate & send aggrement

	return httphelper.HttpSuccessCreated(c, "success invest", nil)
}

// Create visit to check validity of borrower
// (POST /loans/{loanID}/visit)
func (s *Server) CreateLoanVisit(c echo.Context, loanID openapi_types.UUID, params generated.CreateLoanVisitParams) error {
	ctx := c.Request().Context()

	payload := generated.CreateLoanVisitJSONRequestBody{}
	if err := s.bindAndValidatePayload(c, &payload); err != nil {
		return err
	}

	loan, err := s.Repository.GetLoanByID(ctx, loanID)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	// assume there's only can visited once
	if loan.FieldValidatorID != nil {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(errors.New("already visited"), http.StatusUnprocessableEntity))
	}

	if constant.LoanStatus(loan.Status) != constant.LOAN_STATUS_PROPOSED {
		return httphelper.HttpRespError(c, apperror.WrapWithCode(fmt.Errorf("loan status must %s", constant.LOAN_STATUS_PROPOSED), http.StatusUnprocessableEntity))
	}

	visit := &models.Visit{
		ID:         uuid.New(),
		LoanID:     loanID,
		EmployeeID: params.LoginUserId,
		Notes:      payload.Notes,
	}

	for _, proof := range payload.Proof {
		visit.Documents = append(visit.Documents, models.Document{
			ID:            uuid.New(),
			ReferenceID:   visit.ID,
			ReferenceType: "table_visits", // table reference, TODO: update using enum constant
			DocumentType:  "proof_visit",  // TODO: update by enum constant
			FileURL:       proof,
		})
	}

	loginUUID, _ := uuid.Parse(params.LoginUserId.String())
	loan.FieldValidatorID = &loginUUID
	err = s.Repository.CreateVisit(ctx, visit, loan)
	if err != nil {
		return httphelper.HttpRespError(c, err)
	}

	return httphelper.HttpSuccessCreated(c, "success visit", nil)
}
