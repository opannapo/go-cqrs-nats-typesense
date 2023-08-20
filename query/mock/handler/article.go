// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/article.go

// Package handler_mock is a generated GoMock package.
package handler_mock

import (
	reflect "reflect"

	v2 "github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockArticleHandlerImpl is a mock of ArticleHandlerImpl interface.
type MockArticleHandlerImpl struct {
	ctrl     *gomock.Controller
	recorder *MockArticleHandlerImplMockRecorder
}

// MockArticleHandlerImplMockRecorder is the mock recorder for MockArticleHandlerImpl.
type MockArticleHandlerImplMockRecorder struct {
	mock *MockArticleHandlerImpl
}

// NewMockArticleHandlerImpl creates a new mock instance.
func NewMockArticleHandlerImpl(ctrl *gomock.Controller) *MockArticleHandlerImpl {
	mock := &MockArticleHandlerImpl{ctrl: ctrl}
	mock.recorder = &MockArticleHandlerImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleHandlerImpl) EXPECT() *MockArticleHandlerImplMockRecorder {
	return m.recorder
}

// GetArticleByID mocks base method.
func (m *MockArticleHandlerImpl) GetArticleByID(c *v2.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticleByID indicates an expected call of GetArticleByID.
func (mr *MockArticleHandlerImplMockRecorder) GetArticleByID(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByID", reflect.TypeOf((*MockArticleHandlerImpl)(nil).GetArticleByID), c)
}

// Search mocks base method.
func (m *MockArticleHandlerImpl) Search(c *v2.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Search indicates an expected call of Search.
func (mr *MockArticleHandlerImplMockRecorder) Search(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockArticleHandlerImpl)(nil).Search), c)
}
