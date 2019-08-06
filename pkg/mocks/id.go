// Code generated by MockGen. DO NOT EDIT.
// Source: id.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIDGenerator is a mock of IDGenerator interface
type MockIDGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIDGeneratorMockRecorder
}

// MockIDGeneratorMockRecorder is the mock recorder for MockIDGenerator
type MockIDGeneratorMockRecorder struct {
	mock *MockIDGenerator
}

// NewMockIDGenerator creates a new mock instance
func NewMockIDGenerator(ctrl *gomock.Controller) *MockIDGenerator {
	mock := &MockIDGenerator{ctrl: ctrl}
	mock.recorder = &MockIDGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDGenerator) EXPECT() *MockIDGeneratorMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockIDGenerator) New() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockIDGeneratorMockRecorder) New() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockIDGenerator)(nil).New))
}