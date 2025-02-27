package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/crisyantoparulian/loansvc/models"
	apperror "github.com/crisyantoparulian/loansvc/utils/app_error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) getLoanByIDWithDetailSQL(ctx context.Context, id uuid.UUID) (*models.Loan, error) {
	var loan models.Loan

	query := r.DB.Debug().WithContext(ctx).
		Joins("Borrower.User").
		Where("loans.id = ?", id)

		// Apply remaining preloads
	query = query.Preload("Investments.InvestorDetail.UserDetail")

	err := query.First(&loan).Error

	return &loan, err
}

func (r *Repository) getBorrowerByIDWithDetailSQL(ctx context.Context, id uuid.UUID) (*models.Borrower, error) {
	var loan models.Borrower

	query := r.DB.Debug().WithContext(ctx).
		Joins("User").
		Where("borrowers.id = ?", id)

	err := query.First(&loan).Error

	return &loan, err
}

func (r *Repository) saveLoan(ctx context.Context, tx *gorm.DB, loan *models.Loan) error {
	if tx == nil {
		tx = r.DB
	}

	if err := tx.WithContext(ctx).Save(loan).Error; err != nil {
		return apperror.WrapWithCode(fmt.Errorf("failed to save loan: %s", err), http.StatusInternalServerError)
	}

	return nil
}

func (r *Repository) getLoansSQL(ctx context.Context, input GetLoansInput) ([]models.Loan, error) {
	var loans []models.Loan

	query := r.DB.WithContext(ctx)

	query = r.getFilterLoans(query, input).Limit(input.Limit).Offset(input.Offset)

	// Sorting logic
	validSortColumns := map[string]bool{
		"principal_amount":        true,
		"total_investment_amount": true,
		"rate":                    true,
		"roi":                     true,
	}

	if input.SortBy != "" && input.Sort != "" && validSortColumns[input.SortBy] {
		sortOrder := "ASC"
		if input.Sort == "desc" {
			sortOrder = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", input.SortBy, sortOrder))
	}

	// Execute query
	if err := query.Find(&loans).Error; err != nil {
		return nil, err
	}

	return loans, nil
}

func (r *Repository) getFilterLoans(query *gorm.DB, input GetLoansInput) *gorm.DB {
	// Filter by status if provided
	if input.Status != "" {
		query = query.Where("status = ?", input.Status)
	}

	return query
}

func (r *Repository) countLoans(ctx context.Context, input GetLoansInput) (int64, error) {
	var count int64

	query := r.DB.WithContext(ctx).Model(&models.Loan{})

	query = r.getFilterLoans(query, input)

	// Count the total records
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) getLoanByIDSQL(ctx context.Context, id uuid.UUID) (*models.Loan, error) {
	var loan models.Loan

	query := r.DB.Debug().WithContext(ctx).
		Where("loans.id = ?", id)

	err := query.First(&loan).Error

	return &loan, err
}

func (r *Repository) createVisitSQL(ctx context.Context, tx *gorm.DB, visit *models.Visit) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.WithContext(ctx).Create(visit).Error
}

func (r *Repository) createDocumentsSQL(ctx context.Context, tx *gorm.DB, documents []models.Document) error {
	if tx == nil {
		tx = r.DB
	}
	if err := tx.WithContext(ctx).Save(&documents).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) createInvestmentSQL(ctx context.Context, tx *gorm.DB, investment *models.Investment) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.WithContext(ctx).Create(investment).Error
}

func (r *Repository) getLoanInvestmentByInvestorIDSQL(ctx context.Context, loanID, investorID uuid.UUID) ([]models.Investment, error) {
	var loan []models.Investment

	query := r.DB.Debug().WithContext(ctx).
		Where("investor_id = ?", investorID).Where("loan_id", loanID)

	err := query.Find(&loan).Error

	return loan, err
}

func (r *Repository) getVisitsSQL(ctx context.Context, input GetVisitInput) ([]models.Visit, error) {
	var visits []models.Visit

	query := r.DB.Debug().WithContext(ctx)
	if input.LoandID != uuid.Nil {
		query = query.Where("loan_id = ?", input.LoandID)
	}

	query = query.Limit(input.Limit)

	err := query.Find(&visits).Error

	return visits, err
}

func (r *Repository) getEmployeeByIDSQL(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	var employee models.Employee

	query := r.DB.Debug().WithContext(ctx).Where("id = ?", id)

	err := query.First(&employee).Error

	return &employee, err
}

func (r *Repository) createDisbursementSQL(ctx context.Context, tx *gorm.DB, disbursement *models.Disbursement) error {
	if tx == nil {
		tx = r.DB
	}
	return tx.WithContext(ctx).Create(disbursement).Error
}
