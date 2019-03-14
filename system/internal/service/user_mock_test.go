// Code generated by MockGen. DO NOT EDIT.
// Source: system/internal/service/user.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	types "github.com/crusttech/crust/system/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockUserService) With(ctx context.Context) UserService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "With", ctx)
	ret0, _ := ret[0].(UserService)
	return ret0
}

// With indicates an expected call of With
func (mr *MockUserServiceMockRecorder) With(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockUserService)(nil).With), ctx)
}

// FindByUsername mocks base method
func (m *MockUserService) FindByUsername(username string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", username)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername
func (mr *MockUserServiceMockRecorder) FindByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockUserService)(nil).FindByUsername), username)
}

// FindByEmail mocks base method
func (m *MockUserService) FindByEmail(email string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockUserServiceMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserService)(nil).FindByEmail), email)
}

// FindByID mocks base method
func (m *MockUserService) FindByID(id uint64) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockUserServiceMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserService)(nil).FindByID), id)
}

// FindByIDs mocks base method
func (m *MockUserService) FindByIDs(id ...uint64) (types.UserSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range id {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByIDs", varargs...)
	ret0, _ := ret[0].(types.UserSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs
func (mr *MockUserServiceMockRecorder) FindByIDs(id ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockUserService)(nil).FindByIDs), id...)
}

// Find mocks base method
func (m *MockUserService) Find(filter *types.UserFilter) (types.UserSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", filter)
	ret0, _ := ret[0].(types.UserSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUserServiceMockRecorder) Find(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserService)(nil).Find), filter)
}

// FindOrCreate mocks base method
func (m *MockUserService) FindOrCreate(arg0 *types.User) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrCreate", arg0)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrCreate indicates an expected call of FindOrCreate
func (mr *MockUserServiceMockRecorder) FindOrCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrCreate", reflect.TypeOf((*MockUserService)(nil).FindOrCreate), arg0)
}

// Create mocks base method
func (m *MockUserService) Create(input *types.User) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserServiceMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserService)(nil).Create), input)
}

// Update mocks base method
func (m *MockUserService) Update(mod *types.User) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", mod)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUserServiceMockRecorder) Update(mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserService)(nil).Update), mod)
}

// Delete mocks base method
func (m *MockUserService) Delete(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUserServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserService)(nil).Delete), id)
}

// Suspend mocks base method
func (m *MockUserService) Suspend(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspend", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Suspend indicates an expected call of Suspend
func (mr *MockUserServiceMockRecorder) Suspend(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspend", reflect.TypeOf((*MockUserService)(nil).Suspend), id)
}

// Unsuspend mocks base method
func (m *MockUserService) Unsuspend(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsuspend", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsuspend indicates an expected call of Unsuspend
func (mr *MockUserServiceMockRecorder) Unsuspend(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsuspend", reflect.TypeOf((*MockUserService)(nil).Unsuspend), id)
}

// ValidateCredentials mocks base method
func (m *MockUserService) ValidateCredentials(username, password string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCredentials", username, password)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateCredentials indicates an expected call of ValidateCredentials
func (mr *MockUserServiceMockRecorder) ValidateCredentials(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCredentials", reflect.TypeOf((*MockUserService)(nil).ValidateCredentials), username, password)
}
