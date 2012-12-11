// Automatically generated by MockGen. DO NOT EDIT!
// Source: statsdwriter.go

package testsupport

import (
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of StatsdClient interface
type MockStatsdClient struct {
	ctrl     *gomock.Controller
	recorder *_MockStatsdClientRecorder
}

// Recorder for MockStatsdClient (not exported)
type _MockStatsdClientRecorder struct {
	mock *MockStatsdClient
}

func NewMockStatsdClient(ctrl *gomock.Controller) *MockStatsdClient {
	mock := &MockStatsdClient{ctrl: ctrl}
	mock.recorder = &_MockStatsdClientRecorder{mock}
	return mock
}

func (_m *MockStatsdClient) EXPECT() *_MockStatsdClientRecorder {
	return _m.recorder
}

func (_m *MockStatsdClient) IncrementSampledCounter(bucket string, n int, srate float32) {
	_m.ctrl.Call(_m, "IncrementSampledCounter", bucket, n, srate)
}

func (_mr *_MockStatsdClientRecorder) IncrementSampledCounter(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IncrementSampledCounter", arg0, arg1, arg2)
}

func (_m *MockStatsdClient) SendSampledTiming(bucket string, ms int, srate float32) {
	_m.ctrl.Call(_m, "SendSampledTiming", bucket, ms, srate)
}

func (_mr *_MockStatsdClientRecorder) SendSampledTiming(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendSampledTiming", arg0, arg1, arg2)
}