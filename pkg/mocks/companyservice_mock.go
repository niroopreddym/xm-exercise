// Code generated by MockGen. DO NOT EDIT.
// Source: companyservices_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/niroopreddym/xm-exercise/pkg/models"
)

// MockDatabaseServicesIface is a mock of DatabaseServicesIface interface.
type MockDatabaseServicesIface struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseServicesIfaceMockRecorder
}

// MockDatabaseServicesIfaceMockRecorder is the mock recorder for MockDatabaseServicesIface.
type MockDatabaseServicesIfaceMockRecorder struct {
	mock *MockDatabaseServicesIface
}

// NewMockDatabaseServicesIface creates a new mock instance.
func NewMockDatabaseServicesIface(ctrl *gomock.Controller) *MockDatabaseServicesIface {
	mock := &MockDatabaseServicesIface{ctrl: ctrl}
	mock.recorder = &MockDatabaseServicesIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseServicesIface) EXPECT() *MockDatabaseServicesIfaceMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockDatabaseServicesIface) CreateCompany(company *models.Company) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", company)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockDatabaseServicesIfaceMockRecorder) CreateCompany(company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockDatabaseServicesIface)(nil).CreateCompany), company)
}

// DeleteCompanyDetails mocks base method.
func (m *MockDatabaseServicesIface) DeleteCompanyDetails(companyID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompanyDetails", companyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompanyDetails indicates an expected call of DeleteCompanyDetails.
func (mr *MockDatabaseServicesIfaceMockRecorder) DeleteCompanyDetails(companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompanyDetails", reflect.TypeOf((*MockDatabaseServicesIface)(nil).DeleteCompanyDetails), companyID)
}

// GetCompanyDetails mocks base method.
func (m *MockDatabaseServicesIface) GetCompanyDetails(companyID int) (*models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyDetails", companyID)
	ret0, _ := ret[0].(*models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyDetails indicates an expected call of GetCompanyDetails.
func (mr *MockDatabaseServicesIfaceMockRecorder) GetCompanyDetails(companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyDetails", reflect.TypeOf((*MockDatabaseServicesIface)(nil).GetCompanyDetails), companyID)
}

// GetListOfAllCompanies mocks base method.
func (m *MockDatabaseServicesIface) GetListOfAllCompanies() ([]*models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListOfAllCompanies")
	ret0, _ := ret[0].([]*models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListOfAllCompanies indicates an expected call of GetListOfAllCompanies.
func (mr *MockDatabaseServicesIfaceMockRecorder) GetListOfAllCompanies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListOfAllCompanies", reflect.TypeOf((*MockDatabaseServicesIface)(nil).GetListOfAllCompanies))
}

// PutCompanyDetails mocks base method.
func (m *MockDatabaseServicesIface) PutCompanyDetails(companyID int, company *models.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCompanyDetails", companyID, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCompanyDetails indicates an expected call of PutCompanyDetails.
func (mr *MockDatabaseServicesIfaceMockRecorder) PutCompanyDetails(companyID, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCompanyDetails", reflect.TypeOf((*MockDatabaseServicesIface)(nil).PutCompanyDetails), companyID, company)
}