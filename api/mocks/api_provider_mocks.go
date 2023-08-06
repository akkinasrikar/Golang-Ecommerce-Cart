// Code generated by MockGen. DO NOT EDIT.
// Source: api_provider_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	dto "github.com/akkinasrikar/ecommerce-cart/api/dto"
	models "github.com/akkinasrikar/ecommerce-cart/models"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetItems mocks base method.
func (m *MockService) GetItems(ecomCtx context.Context) (dto.ItemsResponse, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ecomCtx)
	ret0, _ := ret[0].(dto.ItemsResponse)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockServiceMockRecorder) GetItems(ecomCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockService)(nil).GetItems), ecomCtx)
}

// MockHttpCall is a mock of HttpCall interface.
type MockHttpCall struct {
	ctrl     *gomock.Controller
	recorder *MockHttpCallMockRecorder
}

// MockHttpCallMockRecorder is the mock recorder for MockHttpCall.
type MockHttpCallMockRecorder struct {
	mock *MockHttpCall
}

// NewMockHttpCall creates a new mock instance.
func NewMockHttpCall(ctrl *gomock.Controller) *MockHttpCall {
	mock := &MockHttpCall{ctrl: ctrl}
	mock.recorder = &MockHttpCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpCall) EXPECT() *MockHttpCallMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHttpCall) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHttpCallMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHttpCall)(nil).Do), req)
}
