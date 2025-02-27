package handler

import (
	"github.com/crisyantoparulian/loansvc/generated"
	"github.com/crisyantoparulian/loansvc/models"
	"github.com/crisyantoparulian/loansvc/repository"
)

// mapping for response detail loan
func mappingToDetailLoanResponse(m *models.Loan) (resp generated.LoanDetail) {
	resp = generated.LoanDetail{
		Id:           m.ID,
		ApprovalDate: m.ApprovalDate,
		Borrower: &generated.Borrower{
			Email:    m.Borrower.User.Email,
			Id:       m.Borrower.ID,
			JoinedAt: m.Borrower.CreatedAt,
			Name:     m.Borrower.User.FullName,
			Phone:    m.Borrower.User.Phone,
			Nik:      m.Borrower.User.NIK,
		},
		FieldValidatorId:      m.FieldValidatorID,
		PrincipalAmount:       m.PrincipalAmount,
		Rate:                  m.Rate,
		Roi:                   m.ROI,
		Status:                generated.LoanDetailStatus(m.Status),
		TotalInvestmentAmount: m.TotalInvestmentAmount,
		TotalInvestors:        m.TotalInvestor,
		CreatedAt:             m.CreatedAt,
	}

	for _, invest := range m.Investments {
		resp.Investments = append(resp.Investments, generated.InvestmentDetail{
			Amount:        invest.Amount,
			InvestmentId:  invest.ID,
			InvestorName:  invest.InvestorDetail.UserDetail.FullName,
			InvestorEmail: invest.InvestorDetail.UserDetail.Email,
			CreatedAt:     invest.CreatedAt,
			InvestorId:    invest.InvestorDetail.ID,
		})
	}

	return
}

// mapping for list response loan
func mappingToLoanListResponse(m []models.Loan, total int64, input repository.GetLoansInput) (resp generated.LoanList) {

	for _, loan := range m {
		resp.Items = append(resp.Items, generated.LoanListDetail{
			ApprovalDate:          loan.ApprovalDate,
			BorrowerId:            loan.BorrowerID,
			CreatedAt:             loan.CreatedAt,
			FieldValidatorId:      loan.FieldValidatorID,
			Id:                    loan.ID,
			PrincipalAmount:       loan.PrincipalAmount,
			Rate:                  loan.Rate,
			Roi:                   loan.ROI,
			Status:                loan.Status,
			TotalInvestmentAmount: loan.TotalInvestmentAmount,
			TotalInvestors:        loan.TotalInvestor,
		})
	}

	resp.Total = total
	resp.Limit = input.Limit
	resp.Offset = input.Offset

	return
}
