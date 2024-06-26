// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChooseCruise/choosecruise-backend/domain (interfaces: LoginUsecase)

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"
	time "time"

	domain "github.com/ChooseCruise/choosecruise-backend/domain"
	tokenutil "github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockLoginUsecase is a mock of LoginUsecase interface.
type MockLoginUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockLoginUsecaseMockRecorder
}

// MockLoginUsecaseMockRecorder is the mock recorder for MockLoginUsecase.
type MockLoginUsecaseMockRecorder struct {
	mock *MockLoginUsecase
}

// NewMockLoginUsecase creates a new mock instance.
func NewMockLoginUsecase(ctrl *gomock.Controller) *MockLoginUsecase {
	mock := &MockLoginUsecase{ctrl: ctrl}
	mock.recorder = &MockLoginUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginUsecase) EXPECT() *MockLoginUsecaseMockRecorder {
	return m.recorder
}

// CreateTokens mocks base method.
func (m *MockLoginUsecase) CreateTokens(arg0 *gin.Context, arg1 *domain.User, arg2, arg3 time.Duration, arg4 tokenutil.Maker) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTokens", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTokens indicates an expected call of CreateTokens.
func (mr *MockLoginUsecaseMockRecorder) CreateTokens(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTokens", reflect.TypeOf((*MockLoginUsecase)(nil).CreateTokens), arg0, arg1, arg2, arg3, arg4)
}

// GetUserByEmail mocks base method.
func (m *MockLoginUsecase) GetUserByEmail(arg0 context.Context, arg1 string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockLoginUsecaseMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockLoginUsecase)(nil).GetUserByEmail), arg0, arg1)
}
