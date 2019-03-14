// Code generated by MockGen. DO NOT EDIT.
// Source: internal/rules/interfaces.go

// Package rules is a generated GoMock package.
package rules

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	factory "github.com/titpetric/factory"
	reflect "reflect"
)

// MockResourcesInterface is a mock of ResourcesInterface interface
type MockResourcesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesInterfaceMockRecorder
}

// MockResourcesInterfaceMockRecorder is the mock recorder for MockResourcesInterface
type MockResourcesInterfaceMockRecorder struct {
	mock *MockResourcesInterface
}

// NewMockResourcesInterface creates a new mock instance
func NewMockResourcesInterface(ctrl *gomock.Controller) *MockResourcesInterface {
	mock := &MockResourcesInterface{ctrl: ctrl}
	mock.recorder = &MockResourcesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourcesInterface) EXPECT() *MockResourcesInterfaceMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockResourcesInterface) With(ctx context.Context, db *factory.DB) ResourcesInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "With", ctx, db)
	ret0, _ := ret[0].(ResourcesInterface)
	return ret0
}

// With indicates an expected call of With
func (mr *MockResourcesInterfaceMockRecorder) With(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockResourcesInterface)(nil).With), ctx, db)
}

// Check mocks base method
func (m *MockResourcesInterface) Check(resource, operation string, fallbacks ...CheckAccessFunc) Access {
	m.ctrl.T.Helper()
	varargs := []interface{}{resource, operation}
	for _, a := range fallbacks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Check", varargs...)
	ret0, _ := ret[0].(Access)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockResourcesInterfaceMockRecorder) Check(resource, operation interface{}, fallbacks ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{resource, operation}, fallbacks...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockResourcesInterface)(nil).Check), varargs...)
}

// Grant mocks base method
func (m *MockResourcesInterface) Grant(roleID uint64, rules []Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Grant", roleID, rules)
	ret0, _ := ret[0].(error)
	return ret0
}

// Grant indicates an expected call of Grant
func (mr *MockResourcesInterfaceMockRecorder) Grant(roleID, rules interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Grant", reflect.TypeOf((*MockResourcesInterface)(nil).Grant), roleID, rules)
}

// Read mocks base method
func (m *MockResourcesInterface) Read(roleID uint64) ([]Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", roleID)
	ret0, _ := ret[0].([]Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockResourcesInterfaceMockRecorder) Read(roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockResourcesInterface)(nil).Read), roleID)
}

// Delete mocks base method
func (m *MockResourcesInterface) Delete(roleID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", roleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockResourcesInterfaceMockRecorder) Delete(roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResourcesInterface)(nil).Delete), roleID)
}
