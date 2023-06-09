// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/event.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "watchAndRun/internal/model"

	gomock "github.com/golang/mock/gomock"
)

// MockEvent is a mock of Event interface.
type MockEvent struct {
	ctrl     *gomock.Controller
	recorder *MockEventMockRecorder
}

// MockEventMockRecorder is the mock recorder for MockEvent.
type MockEventMockRecorder struct {
	mock *MockEvent
}

// NewMockEvent creates a new mock instance.
func NewMockEvent(ctrl *gomock.Controller) *MockEvent {
	mock := &MockEvent{ctrl: ctrl}
	mock.recorder = &MockEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvent) EXPECT() *MockEventMockRecorder {
	return m.recorder
}

// AddEvent mocks base method.
func (m *MockEvent) AddEvent(event model.Event) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEvent", event)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddEvent indicates an expected call of AddEvent.
func (mr *MockEventMockRecorder) AddEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEvent", reflect.TypeOf((*MockEvent)(nil).AddEvent), event)
}

// GetAllEvents mocks base method.
func (m *MockEvent) GetAllEvents() ([]model.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvents")
	ret0, _ := ret[0].([]model.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEvents indicates an expected call of GetAllEvents.
func (mr *MockEventMockRecorder) GetAllEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvents", reflect.TypeOf((*MockEvent)(nil).GetAllEvents))
}
