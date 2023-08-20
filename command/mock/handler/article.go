// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/article.go

// Package handler_mock is a generated GoMock package.
package handler_mock

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
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

// CreateArticle mocks base method.
func (m *MockArticleHandlerImpl) CreateArticle(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockArticleHandlerImplMockRecorder) CreateArticle(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockArticleHandlerImpl)(nil).CreateArticle), c)
}