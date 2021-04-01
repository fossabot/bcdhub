// Code generated by MockGen. DO NOT EDIT.
// Source: tokenbalance/repository.go

// Package mock_tokenbalance is a generated GoMock package.
package tokenbalance

import (
	tb "github.com/baking-bad/bcdhub/internal/models/tokenbalance"
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

// Get mocks base method
func (m *MockRepository) Get(network, contract, address string, tokenID uint64) (tb.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", network, contract, address, tokenID)
	ret0, _ := ret[0].(tb.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(network, contract, address, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), network, contract, address, tokenID)
}

// GetAccountBalances mocks base method
func (m *MockRepository) GetAccountBalances(network, address, contract string, size, offset int64) ([]tb.TokenBalance, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountBalances", network, address, contract, size, offset)
	ret0, _ := ret[0].([]tb.TokenBalance)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAccountBalances indicates an expected call of GetAccountBalances
func (mr *MockRepositoryMockRecorder) GetAccountBalances(network, address, contract, size, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountBalances", reflect.TypeOf((*MockRepository)(nil).GetAccountBalances), network, address, contract, size, offset)
}

// GetHolders mocks base method
func (m *MockRepository) GetHolders(network, contract string, tokenID uint64) ([]tb.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHolders", network, contract, tokenID)
	ret0, _ := ret[0].([]tb.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHolders indicates an expected call of GetHolders
func (mr *MockRepositoryMockRecorder) GetHolders(network, contract, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHolders", reflect.TypeOf((*MockRepository)(nil).GetHolders), network, contract, tokenID)
}

// Batch mocks base method
func (m *MockRepository) Batch(network string, addresses []string) (map[string][]tb.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Batch", network, addresses)
	ret0, _ := ret[0].(map[string][]tb.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batch indicates an expected call of Batch
func (mr *MockRepositoryMockRecorder) Batch(network, addresses interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockRepository)(nil).Batch), network, addresses)
}

// CountByContract mocks base method
func (m *MockRepository) CountByContract(network, address string) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByContract", network, address)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByContract indicates an expected call of CountByContract
func (mr *MockRepositoryMockRecorder) CountByContract(network, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByContract", reflect.TypeOf((*MockRepository)(nil).CountByContract), network, address)
}
