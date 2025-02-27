package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/crisyantoparulian/loansvc/models"
	apperror "github.com/crisyantoparulian/loansvc/utils/app_error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// get loans detail, with borrower & investment detail
func (r *Repository) GetLoanByIDWithDetailSQL(ctx context.Context, id uuid.UUID) (loan *models.Loan, err error) {
	loan, err = r.getLoanByIDWithDetailSQL(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperror.WrapWithCode(fmt.Errorf("loan with ID %s not found", id), http.StatusNotFound)
			return
		}
		err = apperror.WrapWithCode(fmt.Errorf("failed to get loan: %s", err), http.StatusInternalServerError)
		return
	}

	return
}

// creating new loan
func (r *Repository) CreateLoan(ctx context.Context, loan *models.Loan) (err error) {

	err = r.saveLoan(ctx, nil, loan)
	if err != nil {
		return apperror.WrapWithCode(fmt.Errorf("failed to create loan: %s", err), http.StatusInternalServerError)
	}

	return nil
}

// get data borrower with its user relation
func (r *Repository) GetBorrowerByIDWithDetail(ctx context.Context, id uuid.UUID) (borrower *models.Borrower, err error) {
	borrower, err = r.getBorrowerByIDWithDetailSQL(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperror.WrapWithCode(fmt.Errorf("borrower with ID %s not found", id), http.StatusNotFound)
			return
		}
		err = apperror.WrapWithCode(fmt.Errorf("failed to get borrower: %s", err), http.StatusInternalServerError)
		return
	}

	return
}

// get data loans list with its total
func (r *Repository) GetListLoans(ctx context.Context, input GetLoansInput) (loans []models.Loan, total int64, err error) {

	loans, err = r.getLoansSQL(ctx, input)
	if err != nil {
		err = apperror.WrapWithCode(fmt.Errorf("failed to get list loans: %s", err), http.StatusInternalServerError)
		return
	}

	total, err = r.countLoans(ctx, input)
	if err != nil {
		err = apperror.WrapWithCode(fmt.Errorf("failed to get total loans: %s", err), http.StatusInternalServerError)
		return
	}

	return
}

// get loan by id (there's no join for this func)
func (r *Repository) GetLoanByID(ctx context.Context, id uuid.UUID) (loan *models.Loan, err error) {

	loan, err = r.getLoanByIDSQL(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperror.WrapWithCode(fmt.Errorf("loan with ID %s not found", id), http.StatusNotFound)
			return
		}
		err = apperror.WrapWithCode(fmt.Errorf("failed to get loan: %s", err), http.StatusInternalServerError)
		return
	}

	return
}

// create visit with its document
func (r *Repository) CreateVisit(ctx context.Context, visit *models.Visit, loan *models.Loan) (err error) {
	// use transaction
	tx := r.DB.Begin()

	// #1 insert to visit
	err = r.createVisitSQL(ctx, tx, visit)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to insert visit: %s", err), http.StatusInternalServerError)
		return
	}

	// #2 insert to documents
	if len(visit.Documents) > 0 {
		err = r.createDocumentsSQL(ctx, tx, visit.Documents)
		if err != nil {
			tx.Rollback()
			err = apperror.WrapWithCode(fmt.Errorf("failed to insert documents: %s", err), http.StatusInternalServerError)
			return
		}
	}

	err = r.saveLoan(ctx, tx, loan)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to save loan: %s", err), http.StatusInternalServerError)
		return
	}

	// Commit transaction if everything is fine
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return
}

// creating new investment
func (r *Repository) CreateInvestment(ctx context.Context, investment *models.Investment, loan *models.Loan) (err error) {
	// use transaction
	tx := r.DB.Begin()

	// #1 insert to visit
	err = r.createInvestmentSQL(ctx, tx, investment)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to insert investment: %s", err), http.StatusInternalServerError)
		return
	}

	// #2 save the loan
	err = r.saveLoan(ctx, tx, loan)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to save loan: %s", err), http.StatusInternalServerError)
		return
	}

	// Commit transaction if everything is fine
	if err := tx.Commit().Error; err != nil {
		err = apperror.WrapWithCode(fmt.Errorf("failed to commit create investment: %s", err), http.StatusInternalServerError)
		return err
	}

	return
}

// get investment by investor
// TODO: make this func more general
func (r *Repository) GetLoanInvestmentByInvestorID(ctx context.Context, loanID, investorID uuid.UUID) (investments []models.Investment, err error) {
	investments, err = r.getLoanInvestmentByInvestorIDSQL(ctx, loanID, investorID)
	if err != nil {
		err = apperror.WrapWithCode(fmt.Errorf("failed to get investment by investor id: %s", err), http.StatusInternalServerError)
		return
	}

	return investments, err
}

// get data visits
func (r *Repository) GetVisits(ctx context.Context, input GetVisitInput) (visits []models.Visit, err error) {
	visits, err = r.getVisitsSQL(ctx, input)
	if err != nil {
		err = apperror.WrapWithCode(fmt.Errorf("failed to get visits: %s", err), http.StatusInternalServerError)
		return
	}

	return visits, err
}

// approve loan
func (r *Repository) ApproveLoan(ctx context.Context, loan *models.Loan) (err error) {
	// use transaction
	tx := r.DB.Begin()

	// #1 save the loan
	err = r.saveLoan(ctx, tx, loan)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to save loan: %s", err), http.StatusInternalServerError)
		return
	}

	// TODO: create logs here

	// Commit transaction if everything is fine
	if err := tx.Commit().Error; err != nil {
		err = apperror.WrapWithCode(fmt.Errorf("failed to commit create investment: %s", err), http.StatusInternalServerError)
		return err
	}

	return
}

// get employee detail
func (r *Repository) GetEmployeeByID(ctx context.Context, id uuid.UUID) (employee *models.Employee, err error) {
	employee, err = r.getEmployeeByIDSQL(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperror.WrapWithCode(fmt.Errorf("employee with ID %s not found", id), http.StatusNotFound)
			return
		}
		err = apperror.WrapWithCode(fmt.Errorf("failed to get employee: %s", err), http.StatusInternalServerError)
		return
	}

	return
}

// approve loan
func (r *Repository) CreateDisbursement(ctx context.Context,
	disbursement *models.Disbursement,
	documents []models.Document,
	loan *models.Loan) (err error) {

	// use transaction
	tx := r.DB.Begin()

	// #1 create disbursement
	err = r.createDisbursementSQL(ctx, tx, disbursement)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to create disbursement: %s", err), http.StatusInternalServerError)
		return
	}

	// #2 create documents
	err = r.createDocumentsSQL(ctx, tx, documents)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to create documents: %s", err), http.StatusInternalServerError)
		return
	}

	// #3 save the loan
	err = r.saveLoan(ctx, tx, loan)
	if err != nil {
		tx.Rollback()
		err = apperror.WrapWithCode(fmt.Errorf("failed to save loan: %s", err), http.StatusInternalServerError)
		return
	}

	// TODO: create logs here

	// Commit transaction if everything is fine
	if err := tx.Commit().Error; err != nil {
		err = apperror.WrapWithCode(fmt.Errorf("failed to commit create investment: %s", err), http.StatusInternalServerError)
		return err
	}

	return
}
