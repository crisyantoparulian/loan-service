// This file contains types that are used in the repository layer. for  INPUT - OUTPUT
package repository

import (
	"github.com/crisyantoparulian/loansvc/generated"
	"github.com/google/uuid"
)

type GetLoansInput struct {
	Sort, SortBy, Status string
	Limit, Offset        int
}

func (input GetLoansInput) FromParam(params generated.GetLoansParams) GetLoansInput {
	if params.Status != nil {
		input.Status = *params.Status
	}

	if params.S != nil {
		input.Sort = string(*params.S)
	}

	if params.SBy != nil {
		input.SortBy = string(*params.SBy)
	}

	if params.Limit != nil {
		input.Limit = *params.Limit
	}

	// set default limit value
	if input.Limit <= 0 {
		input.Limit = 10
	}

	if params.Offset != nil {
		input.Offset = *params.Offset
	}

	return input
}

type LoanVisitInput struct {
	LoanID uuid.UUID
	Proof  []string
	Note   string
}

type GetInvestmentsInput struct {
	Sort, SortBy, Status string
	Limit, Offset        int
}

type GetVisitInput struct {
	LoandID uuid.UUID
	Limit   int
}
