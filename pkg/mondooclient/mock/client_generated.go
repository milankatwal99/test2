// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mondooclient "go.mondoo.com/mondoo-operator/pkg/mondooclient"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ExchangeRegistrationToken mocks base method.
func (m *MockClient) ExchangeRegistrationToken(arg0 context.Context, arg1 *mondooclient.ExchangeRegistrationTokenInput) (*mondooclient.ExchangeRegistrationTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeRegistrationToken", arg0, arg1)
	ret0, _ := ret[0].(*mondooclient.ExchangeRegistrationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExchangeRegistrationToken indicates an expected call of ExchangeRegistrationToken.
func (mr *MockClientMockRecorder) ExchangeRegistrationToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeRegistrationToken", reflect.TypeOf((*MockClient)(nil).ExchangeRegistrationToken), arg0, arg1)
}

// HealthCheck mocks base method.
func (m *MockClient) HealthCheck(arg0 context.Context, arg1 *mondooclient.HealthCheckRequest) (*mondooclient.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0, arg1)
	ret0, _ := ret[0].(*mondooclient.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockClientMockRecorder) HealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockClient)(nil).HealthCheck), arg0, arg1)
}

// IntegrationCheckIn mocks base method.
func (m *MockClient) IntegrationCheckIn(arg0 context.Context, arg1 *mondooclient.IntegrationCheckInInput) (*mondooclient.IntegrationCheckInOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntegrationCheckIn", arg0, arg1)
	ret0, _ := ret[0].(*mondooclient.IntegrationCheckInOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntegrationCheckIn indicates an expected call of IntegrationCheckIn.
func (mr *MockClientMockRecorder) IntegrationCheckIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntegrationCheckIn", reflect.TypeOf((*MockClient)(nil).IntegrationCheckIn), arg0, arg1)
}

// IntegrationRegister mocks base method.
func (m *MockClient) IntegrationRegister(arg0 context.Context, arg1 *mondooclient.IntegrationRegisterInput) (*mondooclient.IntegrationRegisterOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntegrationRegister", arg0, arg1)
	ret0, _ := ret[0].(*mondooclient.IntegrationRegisterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntegrationRegister indicates an expected call of IntegrationRegister.
func (mr *MockClientMockRecorder) IntegrationRegister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntegrationRegister", reflect.TypeOf((*MockClient)(nil).IntegrationRegister), arg0, arg1)
}

// IntegrationReportStatus mocks base method.
func (m *MockClient) IntegrationReportStatus(arg0 context.Context, arg1 *mondooclient.ReportStatusRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntegrationReportStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IntegrationReportStatus indicates an expected call of IntegrationReportStatus.
func (mr *MockClientMockRecorder) IntegrationReportStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntegrationReportStatus", reflect.TypeOf((*MockClient)(nil).IntegrationReportStatus), arg0, arg1)
}

// RunKubernetesManifest mocks base method.
func (m *MockClient) RunKubernetesManifest(arg0 context.Context, arg1 *mondooclient.KubernetesManifestJob) (*mondooclient.ScanResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunKubernetesManifest", arg0, arg1)
	ret0, _ := ret[0].(*mondooclient.ScanResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunKubernetesManifest indicates an expected call of RunKubernetesManifest.
func (mr *MockClientMockRecorder) RunKubernetesManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunKubernetesManifest", reflect.TypeOf((*MockClient)(nil).RunKubernetesManifest), arg0, arg1)
}

// ScanKubernetesResources mocks base method.
func (m *MockClient) ScanKubernetesResources(ctx context.Context, integrationMrn string, scanContainerImages bool) (*mondooclient.ScanResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanKubernetesResources", ctx, integrationMrn, scanContainerImages)
	ret0, _ := ret[0].(*mondooclient.ScanResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanKubernetesResources indicates an expected call of ScanKubernetesResources.
func (mr *MockClientMockRecorder) ScanKubernetesResources(ctx, integrationMrn, scanContainerImages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanKubernetesResources", reflect.TypeOf((*MockClient)(nil).ScanKubernetesResources), ctx, integrationMrn, scanContainerImages)
}
