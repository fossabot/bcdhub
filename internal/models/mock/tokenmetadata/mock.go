// Code generated by MockGen. DO NOT EDIT.
// Source: internal/models/tokenmetadata/repository.go

// Package tokenmetadata is a generated GoMock package.
package tokenmetadata

import (
	model "github.com/baking-bad/bcdhub/internal/models/tokenmetadata"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
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
func (m *MockRepository) Get(ctx []model.GetContext, size, offset int64) ([]model.TokenMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, size, offset)
	ret0, _ := ret[0].([]model.TokenMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(ctx, size, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, size, offset)
}

// GetAll mocks base method
func (m *MockRepository) GetAll(ctx ...model.GetContext) ([]model.TokenMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAll", varargs...)
	ret0, _ := ret[0].([]model.TokenMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockRepositoryMockRecorder) GetAll(ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), ctx...)
}

// GetOne mocks base method
func (m *MockRepository) GetOne(contract string, tokenID uint64) (*model.TokenMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", contract, tokenID)
	ret0, _ := ret[0].(*model.TokenMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockRepositoryMockRecorder) GetOne(contract, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockRepository)(nil).GetOne), contract, tokenID)
}

// GetRecent mocks base method
func (m *MockRepository) GetRecent(since time.Time, ctx ...model.GetContext) ([]model.TokenMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{since}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecent", varargs...)
	ret0, _ := ret[0].([]model.TokenMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecent indicates an expected call of GetRecent
func (mr *MockRepositoryMockRecorder) GetRecent(since interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{since}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecent", reflect.TypeOf((*MockRepository)(nil).GetRecent), varargs...)
}

// GetWithExtras mocks base method
func (m *MockRepository) GetWithExtras() ([]model.TokenMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithExtras")
	ret0, _ := ret[0].([]model.TokenMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithExtras indicates an expected call of GetWithExtras
func (mr *MockRepositoryMockRecorder) GetWithExtras() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithExtras", reflect.TypeOf((*MockRepository)(nil).GetWithExtras))
}

// Count mocks base method
func (m *MockRepository) Count(ctx []model.GetContext) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockRepositoryMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRepository)(nil).Count), ctx)
}
