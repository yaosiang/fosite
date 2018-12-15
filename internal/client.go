// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: Client)

// Package internal is a generated GoMock package.
package internal

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetAudience mocks base method
func (m *MockClient) GetAudience() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetAudience")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetAudience indicates an expected call of GetAudience
func (mr *MockClientMockRecorder) GetAudience() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAudience", reflect.TypeOf((*MockClient)(nil).GetAudience))
}

// GetGrantTypes mocks base method
func (m *MockClient) GetGrantTypes() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetGrantTypes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetGrantTypes indicates an expected call of GetGrantTypes
func (mr *MockClientMockRecorder) GetGrantTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantTypes", reflect.TypeOf((*MockClient)(nil).GetGrantTypes))
}

// GetHashedSecret mocks base method
func (m *MockClient) GetHashedSecret() []byte {
	ret := m.ctrl.Call(m, "GetHashedSecret")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetHashedSecret indicates an expected call of GetHashedSecret
func (mr *MockClientMockRecorder) GetHashedSecret() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashedSecret", reflect.TypeOf((*MockClient)(nil).GetHashedSecret))
}

// GetID mocks base method
func (m *MockClient) GetID() string {
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockClientMockRecorder) GetID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockClient)(nil).GetID))
}

// GetRedirectURIs mocks base method
func (m *MockClient) GetRedirectURIs() []string {
	ret := m.ctrl.Call(m, "GetRedirectURIs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetRedirectURIs indicates an expected call of GetRedirectURIs
func (mr *MockClientMockRecorder) GetRedirectURIs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedirectURIs", reflect.TypeOf((*MockClient)(nil).GetRedirectURIs))
}

// GetResponseTypes mocks base method
func (m *MockClient) GetResponseTypes() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetResponseTypes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetResponseTypes indicates an expected call of GetResponseTypes
func (mr *MockClientMockRecorder) GetResponseTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseTypes", reflect.TypeOf((*MockClient)(nil).GetResponseTypes))
}

// GetScopes mocks base method
func (m *MockClient) GetScopes() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetScopes indicates an expected call of GetScopes
func (mr *MockClientMockRecorder) GetScopes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScopes", reflect.TypeOf((*MockClient)(nil).GetScopes))
}

// IsPublic mocks base method
func (m *MockClient) IsPublic() bool {
	ret := m.ctrl.Call(m, "IsPublic")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPublic indicates an expected call of IsPublic
func (mr *MockClientMockRecorder) IsPublic() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPublic", reflect.TypeOf((*MockClient)(nil).IsPublic))
}
