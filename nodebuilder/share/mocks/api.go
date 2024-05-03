// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sunrise-zone/sunrise-node/nodebuilder/share (interfaces: Module)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	rsmt2d "github.com/celestiaorg/rsmt2d"
	gomock "github.com/golang/mock/gomock"

	// da "github.com/sunriselayer/sunrise/pkg/da"
	header "github.com/sunrise-zone/sunrise-node/header"
	share "github.com/sunrise-zone/sunrise-node/share"
)

// MockModule is a mock of Module interface.
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule.
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance.
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// GetEDS mocks base method.
func (m *MockModule) GetEDS(arg0 context.Context, arg1 *header.ExtendedHeader) (*rsmt2d.ExtendedDataSquare, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEDS", arg0, arg1)
	ret0, _ := ret[0].(*rsmt2d.ExtendedDataSquare)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEDS indicates an expected call of GetEDS.
func (mr *MockModuleMockRecorder) GetEDS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEDS", reflect.TypeOf((*MockModule)(nil).GetEDS), arg0, arg1)
}

// GetShare mocks base method.
func (m *MockModule) GetShare(arg0 context.Context, arg1 *header.ExtendedHeader, arg2, arg3 int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShare", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShare indicates an expected call of GetShare.
func (mr *MockModuleMockRecorder) GetShare(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShare", reflect.TypeOf((*MockModule)(nil).GetShare), arg0, arg1, arg2, arg3)
}

// GetSharesByNamespace mocks base method.
func (m *MockModule) GetSharesByNamespace(arg0 context.Context, arg1 *header.ExtendedHeader, arg2 share.Namespace) (share.NamespacedShares, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharesByNamespace", arg0, arg1, arg2)
	ret0, _ := ret[0].(share.NamespacedShares)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharesByNamespace indicates an expected call of GetSharesByNamespace.
func (mr *MockModuleMockRecorder) GetSharesByNamespace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharesByNamespace", reflect.TypeOf((*MockModule)(nil).GetSharesByNamespace), arg0, arg1, arg2)
}

// SharesAvailable mocks base method.
func (m *MockModule) SharesAvailable(arg0 context.Context, arg1 *header.ExtendedHeader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SharesAvailable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SharesAvailable indicates an expected call of SharesAvailable.
func (mr *MockModuleMockRecorder) SharesAvailable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SharesAvailable", reflect.TypeOf((*MockModule)(nil).SharesAvailable), arg0, arg1)
}
