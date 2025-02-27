package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Loan struct {
	ID                    uuid.UUID  `gorm:"type:uuid;primaryKey"`
	LoanCode              string     `gorm:"type:char(14);not null"`
	BorrowerID            uuid.UUID  `gorm:"type:uuid;not null"`
	FieldValidatorID      *uuid.UUID `gorm:"type:uuid"`
	PrincipalAmount       float64    `gorm:"not null"`
	TotalInvestmentAmount float64    `gorm:"default:0"`
	TotalInvestor         int        `gorm:"default:0"`

	Rate         float64    `gorm:"not null"`
	ROI          float64    `gorm:"not null"`
	Status       string     `gorm:"type:varchar(50);not null"`
	ApprovalDate *time.Time `gorm:"null"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime"`

	// Relations
	Investments  []Investment  `gorm:"foreignKey:LoanID"`
	Disbursement *Disbursement `gorm:"foreignKey:LoanID"`
	Borrower     Borrower      `gorm:"foreignKey:BorrowerID"`
	Visits       []Visit       `gorm:"foreignKey:LoanID"`
}

type Investment struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	LoanID     uuid.UUID  `gorm:"type:uuid;not null"`
	InvestorID uuid.UUID  `gorm:"type:uuid;not null"`
	Amount     float64    `gorm:"not null"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"type:timestamp"`

	// relations
	InvestorDetail Investor `gorm:"foreignKey:InvestorID"`
}

type Disbursement struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	LoanID             uuid.UUID `gorm:"type:uuid;not null"`
	FieldOfficerID     uuid.UUID `gorm:"type:uuid;not null"`
	DateOfDisbursement time.Time `gorm:"not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
}

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName  string         `gorm:"type:text;not null"`
	Email     string         `gorm:"type:varchar(255);unique;not null"`
	Phone     string         `gorm:"type:varchar(20);unique;not null"`
	NIK       *string        `gorm:"type:char(16);unique"` // Nullable field
	UserRole  pq.StringArray `gorm:"type:text[];not null"` // Array of roles (e.g., "borrower", "investor", etc.)
	CreatedAt time.Time      `gorm:"type:timestamp;default:now()"`
	UpdatedAt *time.Time     `gorm:"type:timestamp"`
}

type Borrower struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	Type      string     `gorm:"type:varchar(255);"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`

	// Relations
	Loans []Loan `gorm:"foreignKey:BorrowerID"`
	User  User   `gorm:"foreignKey:UserID"`
}

type Document struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	ReferenceID   uuid.UUID `gorm:"type:uuid;not null"`
	ReferenceType string    `gorm:"type:varchar(255);not null"` // relation to table
	DocumentType  string    `gorm:"type:varchar(255);not null"` // type of document uploaded, e.g proof_visit
	FileURL       string    `gorm:"type:text;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type Investor struct {
	ID              uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID          uuid.UUID  `gorm:"type:uuid;not null;unique"`
	Company         *string    `gorm:"type:varchar(255)"`
	InvestmentLimit float64    `gorm:"type:decimal(12,2);not null;check:investment_limit >= 0"`
	CreatedAt       time.Time  `gorm:"type:timestamp;default:now()"`
	UpdatedAt       *time.Time `gorm:"type:timestamp"`

	// relation
	UserDetail User `gorm:"foreignKey:UserID"`
}

type Visit struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	LoanID     uuid.UUID  `gorm:"type:uuid;not null;index"`
	EmployeeID uuid.UUID  `gorm:"type:uuid;not null;index"`
	Notes      string     `gorm:"type:text;not null"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime"`

	// Relations
	Documents []Document `gorm:"foreignKey:ReferenceID"`
	Employee  Employee   `gorm:"foreignKey:EmployeeID"`
}

type Employee struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Phone     string    `gorm:"type:varchar(20);not null"`
	Role      string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}
