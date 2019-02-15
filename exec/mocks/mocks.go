// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MiguelMoll/joycast/exec (interfaces: Executor)

// Package mock_exec is a generated GoMock package.
package mock_exec

import (
	exec "github.com/MiguelMoll/joycast/exec"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockExecutor is a mock of Executor interface
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockExecutor) Run(arg0 string, arg1 ...string) (exec.Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(exec.Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockExecutorMockRecorder) Run(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockExecutor)(nil).Run), varargs...)
}