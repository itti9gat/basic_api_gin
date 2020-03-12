// Code generated by MockGen. DO NOT EDIT.
// Source: iiujapp.tech/susaday/repo (interfaces: CategoryINF)

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	gomock "github.com/golang/mock/gomock"
	model "iiujapp.tech/susaday/model"
	reflect "reflect"
)

// MockCategoryINF is a mock of CategoryINF interface
type MockCategoryINF struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryINFMockRecorder
}

// MockCategoryINFMockRecorder is the mock recorder for MockCategoryINF
type MockCategoryINFMockRecorder struct {
	mock *MockCategoryINF
}

// NewMockCategoryINF creates a new mock instance
func NewMockCategoryINF(ctrl *gomock.Controller) *MockCategoryINF {
	mock := &MockCategoryINF{ctrl: ctrl}
	mock.recorder = &MockCategoryINFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategoryINF) EXPECT() *MockCategoryINFMockRecorder {
	return m.recorder
}

// FindCategory mocks base method
func (m *MockCategoryINF) FindCategory(arg0 string) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategory", arg0)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategory indicates an expected call of FindCategory
func (mr *MockCategoryINFMockRecorder) FindCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategory", reflect.TypeOf((*MockCategoryINF)(nil).FindCategory), arg0)
}