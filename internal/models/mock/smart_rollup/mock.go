// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=../mock/smart_rollup/mock.go -package=smart_rollup -typed
//
// Package smart_rollup is a generated GoMock package.
package smart_rollup

import (
	reflect "reflect"

	smartrollup "github.com/baking-bad/bcdhub/internal/models/smart_rollup"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRepository) Get(address string) (smartrollup.SmartRollup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", address)
	ret0, _ := ret[0].(smartrollup.SmartRollup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(address any) *RepositoryGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), address)
	return &RepositoryGetCall{Call: call}
}

// RepositoryGetCall wrap *gomock.Call
type RepositoryGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetCall) Return(arg0 smartrollup.SmartRollup, arg1 error) *RepositoryGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetCall) Do(f func(string) (smartrollup.SmartRollup, error)) *RepositoryGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetCall) DoAndReturn(f func(string) (smartrollup.SmartRollup, error)) *RepositoryGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockRepository) List(limit, offset int64, sort string) ([]smartrollup.SmartRollup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", limit, offset, sort)
	ret0, _ := ret[0].([]smartrollup.SmartRollup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(limit, offset, sort any) *RepositoryListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), limit, offset, sort)
	return &RepositoryListCall{Call: call}
}

// RepositoryListCall wrap *gomock.Call
type RepositoryListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryListCall) Return(arg0 []smartrollup.SmartRollup, arg1 error) *RepositoryListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryListCall) Do(f func(int64, int64, string) ([]smartrollup.SmartRollup, error)) *RepositoryListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryListCall) DoAndReturn(f func(int64, int64, string) ([]smartrollup.SmartRollup, error)) *RepositoryListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
