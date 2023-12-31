// Code generated by MockGen. DO NOT EDIT.
// Source: order.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	domain "grey/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderInterface is a mock of OrderInterface interface.
type MockOrderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOrderInterfaceMockRecorder
}

// MockOrderInterfaceMockRecorder is the mock recorder for MockOrderInterface.
type MockOrderInterfaceMockRecorder struct {
	mock *MockOrderInterface
}

// NewMockOrderInterface creates a new mock instance.
func NewMockOrderInterface(ctrl *gomock.Controller) *MockOrderInterface {
	mock := &MockOrderInterface{ctrl: ctrl}
	mock.recorder = &MockOrderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderInterface) EXPECT() *MockOrderInterfaceMockRecorder {
	return m.recorder
}

// DetailOrder mocks base method.
func (m *MockOrderInterface) DetailOrder(ctx context.Context, userId, orderId int) (domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetailOrder", ctx, userId, orderId)
	ret0, _ := ret[0].(domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetailOrder indicates an expected call of DetailOrder.
func (mr *MockOrderInterfaceMockRecorder) DetailOrder(ctx, userId, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetailOrder", reflect.TypeOf((*MockOrderInterface)(nil).DetailOrder), ctx, userId, orderId)
}

// ListOrder mocks base method.
func (m *MockOrderInterface) ListOrder(ctx context.Context, userId int) (domain.OrderList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrder", ctx, userId)
	ret0, _ := ret[0].(domain.OrderList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrder indicates an expected call of ListOrder.
func (mr *MockOrderInterfaceMockRecorder) ListOrder(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrder", reflect.TypeOf((*MockOrderInterface)(nil).ListOrder), ctx, userId)
}
