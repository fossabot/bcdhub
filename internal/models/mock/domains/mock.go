// Code generated by MockGen. DO NOT EDIT.
// Source: domains/repository.go

// Package mock_domains is a generated GoMock package.
package domains

import (
	domainModel "github.com/baking-bad/bcdhub/internal/models/domains"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// TokenBalances mocks base method
func (m *MockRepository) TokenBalances(network, contract, address string, size, offset int64) (domainModel.TokenBalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenBalances", network, contract, address, size, offset)
	ret0, _ := ret[0].(domainModel.TokenBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenBalances indicates an expected call of TokenBalances
func (mr *MockRepositoryMockRecorder) TokenBalances(network, contract, address, size, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenBalances", reflect.TypeOf((*MockRepository)(nil).TokenBalances), network, contract, address, size, offset)
}
