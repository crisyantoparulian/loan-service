// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	"github.com/crisyantoparulian/loansvc/models"
	"github.com/google/uuid"
)

type RepositoryInterface interface {
	// get data
	GetLoanByIDWithDetailSQL(ctx context.Context, id uuid.UUID) (loan *models.Loan, err error)
	GetBorrowerByIDWithDetail(ctx context.Context, id uuid.UUID) (*models.Borrower, error)
	GetListLoans(ctx context.Context, input GetLoansInput) (loans []models.Loan, total int64, err error)
	GetLoanByID(ctx context.Context, id uuid.UUID) (loan *models.Loan, err error)
	GetLoanInvestmentByInvestorID(ctx context.Context, loanID, investorID uuid.UUID) (investments []models.Investment, err error)
	GetVisits(ctx context.Context, input GetVisitInput) (visits []models.Visit, err error)
	GetEmployeeByID(ctx context.Context, id uuid.UUID) (employee *models.Employee, err error)

	// create or update data
	CreateLoan(ctx context.Context, loan *models.Loan) (err error)
	CreateVisit(ctx context.Context, visit *models.Visit, loan *models.Loan) (err error)
	CreateInvestment(ctx context.Context, investment *models.Investment, loan *models.Loan) (err error)
	ApproveLoan(ctx context.Context, loan *models.Loan) (err error)
	CreateDisbursement(ctx context.Context, disbursement *models.Disbursement, documents []models.Document, loan *models.Loan) (err error)
}
