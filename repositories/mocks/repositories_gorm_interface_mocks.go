// Code generated by MockGen. DO NOT EDIT.
// Source: repositories_gorm_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/akkinasrikar/ecommerce-cart/models"
	entities "github.com/akkinasrikar/ecommerce-cart/models/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// AddToCart mocks base method.
func (m *MockRepositoryInterface) AddToCart(userDetails entities.EcomUsers, Id int) (entities.Item, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", userDetails, Id)
	ret0, _ := ret[0].(entities.Item)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// AddToCart indicates an expected call of AddToCart.
func (mr *MockRepositoryInterfaceMockRecorder) AddToCart(userDetails, Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockRepositoryInterface)(nil).AddToCart), userDetails, Id)
}

// ConsumeKafkaData mocks base method.
func (m *MockRepositoryInterface) ConsumeKafkaData(ctx context.Context, data entities.Consume) (entities.Consume, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeKafkaData", ctx, data)
	ret0, _ := ret[0].(entities.Consume)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// ConsumeKafkaData indicates an expected call of ConsumeKafkaData.
func (mr *MockRepositoryInterfaceMockRecorder) ConsumeKafkaData(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeKafkaData", reflect.TypeOf((*MockRepositoryInterface)(nil).ConsumeKafkaData), ctx, data)
}

// CreateAddress mocks base method.
func (m *MockRepositoryInterface) CreateAddress(addressDetails entities.DeliveryAddress) (entities.DeliveryAddress, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddress", addressDetails)
	ret0, _ := ret[0].(entities.DeliveryAddress)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// CreateAddress indicates an expected call of CreateAddress.
func (mr *MockRepositoryInterfaceMockRecorder) CreateAddress(addressDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateAddress), addressDetails)
}

// CreateCardDetails mocks base method.
func (m *MockRepositoryInterface) CreateCardDetails(cardDetails entities.CardDetails) (entities.CardDetails, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCardDetails", cardDetails)
	ret0, _ := ret[0].(entities.CardDetails)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// CreateCardDetails indicates an expected call of CreateCardDetails.
func (mr *MockRepositoryInterfaceMockRecorder) CreateCardDetails(cardDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCardDetails", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateCardDetails), cardDetails)
}

// CreateEcomAccount mocks base method.
func (m *MockRepositoryInterface) CreateEcomAccount(ecomAccountDetails entities.EcomUsers) (entities.EcomUsers, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEcomAccount", ecomAccountDetails)
	ret0, _ := ret[0].(entities.EcomUsers)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// CreateEcomAccount indicates an expected call of CreateEcomAccount.
func (mr *MockRepositoryInterfaceMockRecorder) CreateEcomAccount(ecomAccountDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEcomAccount", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateEcomAccount), ecomAccountDetails)
}

// CreateOrder mocks base method.
func (m *MockRepositoryInterface) CreateOrder(orderDetails entities.Order) (entities.Order, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", orderDetails)
	ret0, _ := ret[0].(entities.Order)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockRepositoryInterfaceMockRecorder) CreateOrder(orderDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateOrder), orderDetails)
}

// CreateProduct mocks base method.
func (m *MockRepositoryInterface) CreateProduct(item entities.Item) (entities.Item, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", item)
	ret0, _ := ret[0].(entities.Item)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockRepositoryInterfaceMockRecorder) CreateProduct(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateProduct), item)
}

// GetAddress mocks base method.
func (m *MockRepositoryInterface) GetAddress(userDetails entities.EcomUsers) ([]entities.DeliveryAddress, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", userDetails)
	ret0, _ := ret[0].([]entities.DeliveryAddress)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockRepositoryInterfaceMockRecorder) GetAddress(userDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAddress), userDetails)
}

// GetAddressById mocks base method.
func (m *MockRepositoryInterface) GetAddressById(addressId string) (entities.DeliveryAddress, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressById", addressId)
	ret0, _ := ret[0].(entities.DeliveryAddress)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetAddressById indicates an expected call of GetAddressById.
func (mr *MockRepositoryInterfaceMockRecorder) GetAddressById(addressId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAddressById), addressId)
}

// GetAllOrderByUserID mocks base method.
func (m *MockRepositoryInterface) GetAllOrderByUserID(ctx context.Context) ([]entities.Order, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrderByUserID", ctx)
	ret0, _ := ret[0].([]entities.Order)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetAllOrderByUserID indicates an expected call of GetAllOrderByUserID.
func (mr *MockRepositoryInterfaceMockRecorder) GetAllOrderByUserID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrderByUserID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAllOrderByUserID), ctx)
}

// GetAllOrders mocks base method.
func (m *MockRepositoryInterface) GetAllOrders() ([]entities.Order, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrders")
	ret0, _ := ret[0].([]entities.Order)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetAllOrders indicates an expected call of GetAllOrders.
func (mr *MockRepositoryInterfaceMockRecorder) GetAllOrders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrders", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAllOrders))
}

// GetAllProducts mocks base method.
func (m *MockRepositoryInterface) GetAllProducts() ([]entities.Item, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts")
	ret0, _ := ret[0].([]entities.Item)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockRepositoryInterfaceMockRecorder) GetAllProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAllProducts))
}

// GetCardDetails mocks base method.
func (m *MockRepositoryInterface) GetCardDetails(userDetails entities.EcomUsers) ([]entities.CardDetails, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardDetails", userDetails)
	ret0, _ := ret[0].([]entities.CardDetails)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetCardDetails indicates an expected call of GetCardDetails.
func (mr *MockRepositoryInterfaceMockRecorder) GetCardDetails(userDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardDetails", reflect.TypeOf((*MockRepositoryInterface)(nil).GetCardDetails), userDetails)
}

// GetCardDetailsById mocks base method.
func (m *MockRepositoryInterface) GetCardDetailsById(cardId string) (entities.CardDetails, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardDetailsById", cardId)
	ret0, _ := ret[0].(entities.CardDetails)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetCardDetailsById indicates an expected call of GetCardDetailsById.
func (mr *MockRepositoryInterfaceMockRecorder) GetCardDetailsById(cardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardDetailsById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetCardDetailsById), cardId)
}

// GetOrderByID mocks base method.
func (m *MockRepositoryInterface) GetOrderByID(orderId string) (entities.Order, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", orderId)
	ret0, _ := ret[0].(entities.Order)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetOrderByID(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetOrderByID), orderId)
}

// GetProductById mocks base method.
func (m *MockRepositoryInterface) GetProductById(id int) (entities.Item, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", id)
	ret0, _ := ret[0].(entities.Item)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockRepositoryInterfaceMockRecorder) GetProductById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetProductById), id)
}

// GetProductFromCart mocks base method.
func (m *MockRepositoryInterface) GetProductFromCart(itemId int) (entities.Item, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductFromCart", itemId)
	ret0, _ := ret[0].(entities.Item)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetProductFromCart indicates an expected call of GetProductFromCart.
func (mr *MockRepositoryInterfaceMockRecorder) GetProductFromCart(itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductFromCart", reflect.TypeOf((*MockRepositoryInterface)(nil).GetProductFromCart), itemId)
}

// GetUserDetails mocks base method.
func (m *MockRepositoryInterface) GetUserDetails(ctx context.Context) (entities.EcomUsers, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", ctx)
	ret0, _ := ret[0].(entities.EcomUsers)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserDetails(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserDetails), ctx)
}

// GetUserDetailsById mocks base method.
func (m *MockRepositoryInterface) GetUserDetailsById(usersId int64) (entities.EcomUsers, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetailsById", usersId)
	ret0, _ := ret[0].(entities.EcomUsers)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// GetUserDetailsById indicates an expected call of GetUserDetailsById.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserDetailsById(usersId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetailsById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserDetailsById), usersId)
}

// Login mocks base method.
func (m *MockRepositoryInterface) Login(userDetails entities.Login) (entities.SignUp, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", userDetails)
	ret0, _ := ret[0].(entities.SignUp)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockRepositoryInterfaceMockRecorder) Login(userDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRepositoryInterface)(nil).Login), userDetails)
}

// SignUp mocks base method.
func (m *MockRepositoryInterface) SignUp(userDetails entities.SignUp) (entities.SignUp, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", userDetails)
	ret0, _ := ret[0].(entities.SignUp)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockRepositoryInterfaceMockRecorder) SignUp(userDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockRepositoryInterface)(nil).SignUp), userDetails)
}

// UpdateEcomAccount mocks base method.
func (m *MockRepositoryInterface) UpdateEcomAccount(ecomAccountDetails entities.EcomUsers, ecomId string) (entities.EcomUsers, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEcomAccount", ecomAccountDetails, ecomId)
	ret0, _ := ret[0].(entities.EcomUsers)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// UpdateEcomAccount indicates an expected call of UpdateEcomAccount.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateEcomAccount(ecomAccountDetails, ecomId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEcomAccount", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateEcomAccount), ecomAccountDetails, ecomId)
}

// UpdateOrderByID mocks base method.
func (m *MockRepositoryInterface) UpdateOrderByID(orderId string, orderDetails entities.Order) (entities.Order, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderByID", orderId, orderDetails)
	ret0, _ := ret[0].(entities.Order)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// UpdateOrderByID indicates an expected call of UpdateOrderByID.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateOrderByID(orderId, orderDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderByID", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateOrderByID), orderId, orderDetails)
}

// UpdateProductByID mocks base method.
func (m *MockRepositoryInterface) UpdateProductByID(id int, item entities.Item) (entities.Item, models.EcomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductByID", id, item)
	ret0, _ := ret[0].(entities.Item)
	ret1, _ := ret[1].(models.EcomError)
	return ret0, ret1
}

// UpdateProductByID indicates an expected call of UpdateProductByID.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateProductByID(id, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductByID", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateProductByID), id, item)
}
