// Code generated by MockGen. DO NOT EDIT.
// Source: notifications/push/fcm/client.go

// Package fcm is a generated GoMock package.
package fcm

import (
	go_fcm "github.com/NaySoftware/go-fcm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFirebaseClient is a mock of FirebaseClient interface
type MockFirebaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseClientMockRecorder
}

// MockFirebaseClientMockRecorder is the mock recorder for MockFirebaseClient
type MockFirebaseClientMockRecorder struct {
	mock *MockFirebaseClient
}

// NewMockFirebaseClient creates a new mock instance
func NewMockFirebaseClient(ctrl *gomock.Controller) *MockFirebaseClient {
	mock := &MockFirebaseClient{ctrl: ctrl}
	mock.recorder = &MockFirebaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFirebaseClient) EXPECT() *MockFirebaseClientMockRecorder {
	return m.recorder
}

// NewFcmRegIdsMsg mocks base method
func (m *MockFirebaseClient) NewFcmRegIdsMsg(tokens []string, body interface{}) *go_fcm.FcmClient {
	ret := m.ctrl.Call(m, "NewFcmRegIdsMsg", tokens, body)
	ret0, _ := ret[0].(*go_fcm.FcmClient)
	return ret0
}

// NewFcmRegIdsMsg indicates an expected call of NewFcmRegIdsMsg
func (mr *MockFirebaseClientMockRecorder) NewFcmRegIdsMsg(tokens, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFcmRegIdsMsg", reflect.TypeOf((*MockFirebaseClient)(nil).NewFcmRegIdsMsg), tokens, body)
}

// Send mocks base method
func (m *MockFirebaseClient) Send() (*go_fcm.FcmResponseStatus, error) {
	ret := m.ctrl.Call(m, "Send")
	ret0, _ := ret[0].(*go_fcm.FcmResponseStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockFirebaseClientMockRecorder) Send() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockFirebaseClient)(nil).Send))
}

// SetNotificationPayload mocks base method
func (m *MockFirebaseClient) SetNotificationPayload(payload *go_fcm.NotificationPayload) *go_fcm.FcmClient {
	ret := m.ctrl.Call(m, "SetNotificationPayload", payload)
	ret0, _ := ret[0].(*go_fcm.FcmClient)
	return ret0
}

// SetNotificationPayload indicates an expected call of SetNotificationPayload
func (mr *MockFirebaseClientMockRecorder) SetNotificationPayload(payload interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNotificationPayload", reflect.TypeOf((*MockFirebaseClient)(nil).SetNotificationPayload), payload)
}
