// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=repository/interfaces.go -destination=repository/interfaces.mock.gen.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	models "github.com/crisyantoparulian/loansvc/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// ApproveLoan mocks base method.
func (m *MockRepositoryInterface) ApproveLoan(ctx context.Context, loan *models.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveLoan", ctx, loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveLoan indicates an expected call of ApproveLoan.
func (mr *MockRepositoryInterfaceMockRecorder) ApproveLoan(ctx, loan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveLoan", reflect.TypeOf((*MockRepositoryInterface)(nil).ApproveLoan), ctx, loan)
}

// CreateDisbursement mocks base method.
func (m *MockRepositoryInterface) CreateDisbursement(ctx context.Context, disbursement *models.Disbursement, documents []models.Document, loan *models.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDisbursement", ctx, disbursement, documents, loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDisbursement indicates an expected call of CreateDisbursement.
func (mr *MockRepositoryInterfaceMockRecorder) CreateDisbursement(ctx, disbursement, documents, loan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDisbursement", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateDisbursement), ctx, disbursement, documents, loan)
}

// CreateInvestment mocks base method.
func (m *MockRepositoryInterface) CreateInvestment(ctx context.Context, investment *models.Investment, loan *models.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvestment", ctx, investment, loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInvestment indicates an expected call of CreateInvestment.
func (mr *MockRepositoryInterfaceMockRecorder) CreateInvestment(ctx, investment, loan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvestment", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateInvestment), ctx, investment, loan)
}

// CreateLoan mocks base method.
func (m *MockRepositoryInterface) CreateLoan(ctx context.Context, loan *models.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoan", ctx, loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLoan indicates an expected call of CreateLoan.
func (mr *MockRepositoryInterfaceMockRecorder) CreateLoan(ctx, loan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoan", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateLoan), ctx, loan)
}

// CreateVisit mocks base method.
func (m *MockRepositoryInterface) CreateVisit(ctx context.Context, visit *models.Visit, loan *models.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVisit", ctx, visit, loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVisit indicates an expected call of CreateVisit.
func (mr *MockRepositoryInterfaceMockRecorder) CreateVisit(ctx, visit, loan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVisit", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateVisit), ctx, visit, loan)
}

// GetBorrowerByIDWithDetail mocks base method.
func (m *MockRepositoryInterface) GetBorrowerByIDWithDetail(ctx context.Context, id uuid.UUID) (*models.Borrower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBorrowerByIDWithDetail", ctx, id)
	ret0, _ := ret[0].(*models.Borrower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBorrowerByIDWithDetail indicates an expected call of GetBorrowerByIDWithDetail.
func (mr *MockRepositoryInterfaceMockRecorder) GetBorrowerByIDWithDetail(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBorrowerByIDWithDetail", reflect.TypeOf((*MockRepositoryInterface)(nil).GetBorrowerByIDWithDetail), ctx, id)
}

// GetEmployeeByID mocks base method.
func (m *MockRepositoryInterface) GetEmployeeByID(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByID", ctx, id)
	ret0, _ := ret[0].(*models.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByID indicates an expected call of GetEmployeeByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetEmployeeByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetEmployeeByID), ctx, id)
}

// GetListLoans mocks base method.
func (m *MockRepositoryInterface) GetListLoans(ctx context.Context, input GetLoansInput) ([]models.Loan, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListLoans", ctx, input)
	ret0, _ := ret[0].([]models.Loan)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetListLoans indicates an expected call of GetListLoans.
func (mr *MockRepositoryInterfaceMockRecorder) GetListLoans(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListLoans", reflect.TypeOf((*MockRepositoryInterface)(nil).GetListLoans), ctx, input)
}

// GetLoanByID mocks base method.
func (m *MockRepositoryInterface) GetLoanByID(ctx context.Context, id uuid.UUID) (*models.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanByID", ctx, id)
	ret0, _ := ret[0].(*models.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanByID indicates an expected call of GetLoanByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetLoanByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetLoanByID), ctx, id)
}

// GetLoanByIDWithDetailSQL mocks base method.
func (m *MockRepositoryInterface) GetLoanByIDWithDetailSQL(ctx context.Context, id uuid.UUID) (*models.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanByIDWithDetailSQL", ctx, id)
	ret0, _ := ret[0].(*models.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanByIDWithDetailSQL indicates an expected call of GetLoanByIDWithDetailSQL.
func (mr *MockRepositoryInterfaceMockRecorder) GetLoanByIDWithDetailSQL(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanByIDWithDetailSQL", reflect.TypeOf((*MockRepositoryInterface)(nil).GetLoanByIDWithDetailSQL), ctx, id)
}

// GetLoanInvestmentByInvestorID mocks base method.
func (m *MockRepositoryInterface) GetLoanInvestmentByInvestorID(ctx context.Context, loanID, investorID uuid.UUID) ([]models.Investment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoanInvestmentByInvestorID", ctx, loanID, investorID)
	ret0, _ := ret[0].([]models.Investment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoanInvestmentByInvestorID indicates an expected call of GetLoanInvestmentByInvestorID.
func (mr *MockRepositoryInterfaceMockRecorder) GetLoanInvestmentByInvestorID(ctx, loanID, investorID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoanInvestmentByInvestorID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetLoanInvestmentByInvestorID), ctx, loanID, investorID)
}

// GetVisits mocks base method.
func (m *MockRepositoryInterface) GetVisits(ctx context.Context, input GetVisitInput) ([]models.Visit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVisits", ctx, input)
	ret0, _ := ret[0].([]models.Visit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVisits indicates an expected call of GetVisits.
func (mr *MockRepositoryInterfaceMockRecorder) GetVisits(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVisits", reflect.TypeOf((*MockRepositoryInterface)(nil).GetVisits), ctx, input)
}
