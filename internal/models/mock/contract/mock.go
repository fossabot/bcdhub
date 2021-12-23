// Code generated by MockGen. DO NOT EDIT.
// Source: internal/models/contract/repository.go

// Package contract is a generated GoMock package.
package contract

import (
	model "github.com/baking-bad/bcdhub/internal/models/contract"
	types "github.com/baking-bad/bcdhub/internal/models/types"
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
func (m *MockRepository) Get(network types.Network, address string) (model.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", network, address)
	ret0, _ := ret[0].(model.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(network, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), network, address)
}

// GetMany mocks base method
func (m *MockRepository) GetMany(by map[string]interface{}) ([]model.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", by)
	ret0, _ := ret[0].([]model.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany
func (mr *MockRepositoryMockRecorder) GetMany(by interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockRepository)(nil).GetMany), by)
}

// GetRandom mocks base method
func (m *MockRepository) GetRandom(network types.Network) (model.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandom", network)
	ret0, _ := ret[0].(model.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandom indicates an expected call of GetRandom
func (mr *MockRepositoryMockRecorder) GetRandom(network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandom", reflect.TypeOf((*MockRepository)(nil).GetRandom), network)
}

// GetTokens mocks base method
func (m *MockRepository) GetTokens(network types.Network, tokenInterface string, offset, size int64) ([]model.Contract, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokens", network, tokenInterface, offset, size)
	ret0, _ := ret[0].([]model.Contract)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTokens indicates an expected call of GetTokens
func (mr *MockRepositoryMockRecorder) GetTokens(network, tokenInterface, offset, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokens", reflect.TypeOf((*MockRepository)(nil).GetTokens), network, tokenInterface, offset, size)
}

// GetSameContracts mocks base method
func (m *MockRepository) GetSameContracts(contact model.Contract, manager string, size, offset int64) (model.SameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSameContracts", contact, manager, size, offset)
	ret0, _ := ret[0].(model.SameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSameContracts indicates an expected call of GetSameContracts
func (mr *MockRepositoryMockRecorder) GetSameContracts(contact, manager, size, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSameContracts", reflect.TypeOf((*MockRepository)(nil).GetSameContracts), contact, manager, size, offset)
}

// GetSimilarContracts mocks base method
func (m *MockRepository) GetSimilarContracts(arg0 model.Contract, arg1, arg2 int64) ([]model.Similar, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSimilarContracts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Similar)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSimilarContracts indicates an expected call of GetSimilarContracts
func (mr *MockRepositoryMockRecorder) GetSimilarContracts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimilarContracts", reflect.TypeOf((*MockRepository)(nil).GetSimilarContracts), arg0, arg1, arg2)
}

// Stats mocks base method
func (m *MockRepository) Stats(c model.Contract) (model.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats", c)
	ret0, _ := ret[0].(model.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockRepositoryMockRecorder) Stats(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockRepository)(nil).Stats), c)
}

// Script mocks base method
func (m *MockRepository) Script(network types.Network, address, symLink string) (model.Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Script", network, address, symLink)
	ret0, _ := ret[0].(model.Script)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Script indicates an expected call of Script
func (mr *MockRepositoryMockRecorder) Script(network, address, symLink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Script", reflect.TypeOf((*MockRepository)(nil).Script), network, address, symLink)
}

// ScriptPart mocks base method
func (m *MockRepository) ScriptPart(network types.Network, address, symLink, part string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptPart", network, address, symLink, part)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScriptPart indicates an expected call of ScriptPart
func (mr *MockRepositoryMockRecorder) ScriptPart(network, address, symLink, part interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptPart", reflect.TypeOf((*MockRepository)(nil).ScriptPart), network, address, symLink, part)
}

// MockScriptRepository is a mock of ScriptRepository interface
type MockScriptRepository struct {
	ctrl     *gomock.Controller
	recorder *MockScriptRepositoryMockRecorder
}

// MockScriptRepositoryMockRecorder is the mock recorder for MockScriptRepository
type MockScriptRepositoryMockRecorder struct {
	mock *MockScriptRepository
}

// NewMockScriptRepository creates a new mock instance
func NewMockScriptRepository(ctrl *gomock.Controller) *MockScriptRepository {
	mock := &MockScriptRepository{ctrl: ctrl}
	mock.recorder = &MockScriptRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScriptRepository) EXPECT() *MockScriptRepositoryMockRecorder {
	return m.recorder
}

// GetScripts mocks base method
func (m *MockScriptRepository) GetScripts(limit, offset int) ([]model.Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScripts", limit, offset)
	ret0, _ := ret[0].([]model.Script)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScripts indicates an expected call of GetScripts
func (mr *MockScriptRepositoryMockRecorder) GetScripts(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScripts", reflect.TypeOf((*MockScriptRepository)(nil).GetScripts), limit, offset)
}

// ByHash mocks base method
func (m *MockScriptRepository) ByHash(hash string) (model.Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByHash", hash)
	ret0, _ := ret[0].(model.Script)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByHash indicates an expected call of ByHash
func (mr *MockScriptRepositoryMockRecorder) ByHash(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByHash", reflect.TypeOf((*MockScriptRepository)(nil).ByHash), hash)
}

// UpdateProjectID mocks base method
func (m *MockScriptRepository) UpdateProjectID(script []model.Script) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectID", script)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectID indicates an expected call of UpdateProjectID
func (mr *MockScriptRepositoryMockRecorder) UpdateProjectID(script interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectID", reflect.TypeOf((*MockScriptRepository)(nil).UpdateProjectID), script)
}

// Code mocks base method
func (m *MockScriptRepository) Code(id int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Code", id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Code indicates an expected call of Code
func (mr *MockScriptRepositoryMockRecorder) Code(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Code", reflect.TypeOf((*MockScriptRepository)(nil).Code), id)
}

// Parameter mocks base method
func (m *MockScriptRepository) Parameter(id int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameter", id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parameter indicates an expected call of Parameter
func (mr *MockScriptRepositoryMockRecorder) Parameter(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameter", reflect.TypeOf((*MockScriptRepository)(nil).Parameter), id)
}

// Storage mocks base method
func (m *MockScriptRepository) Storage(id int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage", id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Storage indicates an expected call of Storage
func (mr *MockScriptRepositoryMockRecorder) Storage(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockScriptRepository)(nil).Storage), id)
}
