// Code generated by MockGen. DO NOT EDIT.
// Source: https://github.com/repofuel/repofuel-ingest/pkg/providers (interfaces: SourceIntegration)

// Package mock_providers is a generated GoMock package.
package mock_providers

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	common "github.com/repofuel/repofuel-go-pkgs/common"
	credentials "github.com/repofuel/repofuel-go-pkgs/credentials"
	identifier "https://github.com/repofuel/repofuel-ingest/pkg/identifier"
	providers "https://github.com/repofuel/repofuel-ingest/pkg/providers"
	status "https://github.com/repofuel/repofuel-ingest/pkg/status"
	reflect "reflect"
)

// MockSourceIntegration is a mock of SourceIntegration interface
type MockSourceIntegration struct {
	ctrl     *gomock.Controller
	recorder *MockSourceIntegrationMockRecorder
}

// MockSourceIntegrationMockRecorder is the mock recorder for MockSourceIntegration
type MockSourceIntegrationMockRecorder struct {
	mock *MockSourceIntegration
}

// NewMockSourceIntegration creates a new mock instance
func NewMockSourceIntegration(ctrl *gomock.Controller) *MockSourceIntegration {
	mock := &MockSourceIntegration{ctrl: ctrl}
	mock.recorder = &MockSourceIntegrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceIntegration) EXPECT() *MockSourceIntegrationMockRecorder {
	return m.recorder
}

// BasicAuth mocks base method
func (m *MockSourceIntegration) BasicAuth(arg0 context.Context) (*credentials.BasicAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BasicAuth", arg0)
	ret0, _ := ret[0].(*credentials.BasicAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BasicAuth indicates an expected call of BasicAuth
func (mr *MockSourceIntegrationMockRecorder) BasicAuth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BasicAuth", reflect.TypeOf((*MockSourceIntegration)(nil).BasicAuth), arg0)
}

// FetchCollaborators mocks base method
func (m *MockSourceIntegration) FetchCollaborators(arg0 context.Context) (map[string]common.Permissions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCollaborators", arg0)
	ret0, _ := ret[0].(map[string]common.Permissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCollaborators indicates an expected call of FetchCollaborators
func (mr *MockSourceIntegrationMockRecorder) FetchCollaborators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCollaborators", reflect.TypeOf((*MockSourceIntegration)(nil).FetchCollaborators), arg0)
}

// FetchRepositoryCollaborators mocks base method
func (m *MockSourceIntegration) FetchRepositoryCollaborators(arg0 context.Context) (map[string]common.Permissions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRepositoryCollaborators", arg0)
	ret0, _ := ret[0].(map[string]common.Permissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRepositoryCollaborators indicates an expected call of FetchRepositoryCollaborators
func (mr *MockSourceIntegrationMockRecorder) FetchRepositoryCollaborators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRepositoryCollaborators", reflect.TypeOf((*MockSourceIntegration)(nil).FetchRepositoryCollaborators), arg0)
}

// FetchRepositoryInfo mocks base method
func (m *MockSourceIntegration) FetchRepositoryInfo(arg0 context.Context) (*common.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRepositoryInfo", arg0)
	ret0, _ := ret[0].(*common.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRepositoryInfo indicates an expected call of FetchRepositoryInfo
func (mr *MockSourceIntegrationMockRecorder) FetchRepositoryInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRepositoryInfo", reflect.TypeOf((*MockSourceIntegration)(nil).FetchRepositoryInfo), arg0)
}

// FinishCheckRun mocks base method
func (m *MockSourceIntegration) FinishCheckRun(arg0 context.Context, arg1 map[string]interface{}, arg2 status.Stage, arg3 providers.CheckRunSummarizer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishCheckRun", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishCheckRun indicates an expected call of FinishCheckRun
func (mr *MockSourceIntegrationMockRecorder) FinishCheckRun(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishCheckRun", reflect.TypeOf((*MockSourceIntegration)(nil).FinishCheckRun), arg0, arg1, arg2, arg3)
}

// ListOpenPullRequests mocks base method
func (m *MockSourceIntegration) ListOpenPullRequests(arg0 context.Context) common.PullRequestItr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenPullRequests", arg0)
	ret0, _ := ret[0].(common.PullRequestItr)
	return ret0
}

// ListOpenPullRequests indicates an expected call of ListOpenPullRequests
func (mr *MockSourceIntegrationMockRecorder) ListOpenPullRequests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenPullRequests", reflect.TypeOf((*MockSourceIntegration)(nil).ListOpenPullRequests), arg0)
}

// StartCheckRun mocks base method
func (m *MockSourceIntegration) StartCheckRun(arg0 context.Context, arg1 identifier.JobID, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCheckRun", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartCheckRun indicates an expected call of StartCheckRun
func (mr *MockSourceIntegrationMockRecorder) StartCheckRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCheckRun", reflect.TypeOf((*MockSourceIntegration)(nil).StartCheckRun), arg0, arg1, arg2)
}
