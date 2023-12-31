// Code generated by MockGen. DO NOT EDIT.
// Source: login_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/akkinasrikar/ecommerce-cart/models"
	responses "github.com/akkinasrikar/ecommerce-cart/models/responses"
	gomock "github.com/golang/mock/gomock"
)

// MockLoginService is a mock of LoginService interface.
type MockLoginService struct {
	ctrl     *gomock.Controller
	recorder *MockLoginServiceMockRecorder
}

// MockLoginServiceMockRecorder is the mock recorder for MockLoginService.
type MockLoginServiceMockRecorder struct {
	mock *MockLoginService
}

// NewMockLoginService creates a new mock instance.
func NewMockLoginService(ctrl *gomock.Controller) *MockLoginService {
	mock := &MockLoginService{ctrl: ctrl}
	mock.recorder = &MockLoginServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginService) EXPECT() *MockLoginServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockLoginService) Login(req models.Login) (responses.Login, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", req)
	ret0, _ := ret[0].(responses.Login)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockLoginServiceMockRecorder) Login(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLoginService)(nil).Login), req)
}

// SignUp mocks base method.
func (m *MockLoginService) SignUp(req models.SignUp) (responses.SingUp, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", req)
	ret0, _ := ret[0].(responses.SingUp)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockLoginServiceMockRecorder) SignUp(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockLoginService)(nil).SignUp), req)
}
