// Code generated by MockGen. DO NOT EDIT.
// Source: git.horahora.org/otoman/user-service.git/protocol (interfaces: UserServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	proto "git.horahora.org/otoman/user-service.git/protocol"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// GetUserForForeignUID mocks base method.
func (m *MockUserServiceClient) GetUserForForeignUID(arg0 context.Context, arg1 *proto.GetForeignUserRequest, arg2 ...grpc.CallOption) (*proto.GetForeignUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserForForeignUID", varargs...)
	ret0, _ := ret[0].(*proto.GetForeignUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserForForeignUID indicates an expected call of GetUserForForeignUID.
func (mr *MockUserServiceClientMockRecorder) GetUserForForeignUID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserForForeignUID", reflect.TypeOf((*MockUserServiceClient)(nil).GetUserForForeignUID), varargs...)
}

// GetUserFromID mocks base method.
func (m *MockUserServiceClient) GetUserFromID(arg0 context.Context, arg1 *proto.GetUserFromIDRequest, arg2 ...grpc.CallOption) (*proto.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserFromID", varargs...)
	ret0, _ := ret[0].(*proto.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFromID indicates an expected call of GetUserFromID.
func (mr *MockUserServiceClientMockRecorder) GetUserFromID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFromID", reflect.TypeOf((*MockUserServiceClient)(nil).GetUserFromID), varargs...)
}

// Login mocks base method.
func (m *MockUserServiceClient) Login(arg0 context.Context, arg1 *proto.LoginRequest, arg2 ...grpc.CallOption) (*proto.LoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*proto.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceClientMockRecorder) Login(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserServiceClient)(nil).Login), varargs...)
}

// Register mocks base method.
func (m *MockUserServiceClient) Register(arg0 context.Context, arg1 *proto.RegisterRequest, arg2 ...grpc.CallOption) (*proto.RegisterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*proto.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceClientMockRecorder) Register(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserServiceClient)(nil).Register), varargs...)
}

// ValidateJWT mocks base method.
func (m *MockUserServiceClient) ValidateJWT(arg0 context.Context, arg1 *proto.ValidateJWTRequest, arg2 ...grpc.CallOption) (*proto.ValidateJWTResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateJWT", varargs...)
	ret0, _ := ret[0].(*proto.ValidateJWTResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateJWT indicates an expected call of ValidateJWT.
func (mr *MockUserServiceClientMockRecorder) ValidateJWT(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateJWT", reflect.TypeOf((*MockUserServiceClient)(nil).ValidateJWT), varargs...)
}
