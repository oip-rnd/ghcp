// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/ghcp/usecases/interfaces (interfaces: CopyUseCase)

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/int128/ghcp/usecases/interfaces"
	reflect "reflect"
)

// MockCopyUseCase is a mock of CopyUseCase interface
type MockCopyUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCopyUseCaseMockRecorder
}

// MockCopyUseCaseMockRecorder is the mock recorder for MockCopyUseCase
type MockCopyUseCaseMockRecorder struct {
	mock *MockCopyUseCase
}

// NewMockCopyUseCase creates a new mock instance
func NewMockCopyUseCase(ctrl *gomock.Controller) *MockCopyUseCase {
	mock := &MockCopyUseCase{ctrl: ctrl}
	mock.recorder = &MockCopyUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCopyUseCase) EXPECT() *MockCopyUseCaseMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockCopyUseCase) Do(arg0 context.Context, arg1 interfaces.CopyUseCaseIn) error {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockCopyUseCaseMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockCopyUseCase)(nil).Do), arg0, arg1)
}
