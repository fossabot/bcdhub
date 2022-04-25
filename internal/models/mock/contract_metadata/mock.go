// Code generated by MockGen. DO NOT EDIT.
// Source: internal/models/contract_metadata/repository.go

// Package contract_metadata is a generated GoMock package.
package contract_metadata

import (
	model "github.com/baking-bad/bcdhub/internal/models/contract_metadata"
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
func (m *MockRepository) Get(address string) (*model.ContractMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", address)
	ret0, _ := ret[0].(*model.ContractMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), address)
}

// GetWithEvents mocks base method
func (m *MockRepository) GetWithEvents(updatedAt uint64) ([]model.ContractMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithEvents", updatedAt)
	ret0, _ := ret[0].([]model.ContractMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithEvents indicates an expected call of GetWithEvents
func (mr *MockRepositoryMockRecorder) GetWithEvents(updatedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithEvents", reflect.TypeOf((*MockRepository)(nil).GetWithEvents), updatedAt)
}

// GetBySlug mocks base method
func (m *MockRepository) GetBySlug(slug string) (*model.ContractMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", slug)
	ret0, _ := ret[0].(*model.ContractMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug
func (mr *MockRepositoryMockRecorder) GetBySlug(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockRepository)(nil).GetBySlug), slug)
}

// GetAliases mocks base method
func (m *MockRepository) GetAliases() ([]model.ContractMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAliases")
	ret0, _ := ret[0].([]model.ContractMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAliases indicates an expected call of GetAliases
func (mr *MockRepositoryMockRecorder) GetAliases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAliases", reflect.TypeOf((*MockRepository)(nil).GetAliases))
}

// Events mocks base method
func (m *MockRepository) Events(address string) (model.Events, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events", address)
	ret0, _ := ret[0].(model.Events)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Events indicates an expected call of Events
func (mr *MockRepositoryMockRecorder) Events(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockRepository)(nil).Events), address)
}
