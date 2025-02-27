package constant

type LoanStatus string

const (
	LOAN_STATUS_PROPOSED  LoanStatus = LoanStatus("proposed")
	LOAN_STATUS_APPROVED  LoanStatus = LoanStatus("approved")
	LOAN_STATUS_INVESTED  LoanStatus = LoanStatus("invested")
	LOAN_STATUS_DISBURSED LoanStatus = LoanStatus("disbursed")
)

// RoleSet for fast validation lookup
var validLoanStatus = map[LoanStatus]bool{
	LOAN_STATUS_PROPOSED:  true,
	LOAN_STATUS_APPROVED:  true,
	LOAN_STATUS_INVESTED:  true,
	LOAN_STATUS_DISBURSED: true,
}

// IsValid checks if a given role is valid
func (r LoanStatus) IsValid() bool {
	_, exists := validLoanStatus[r]
	return exists
}

func (r LoanStatus) String() string {
	return string(r)
}
