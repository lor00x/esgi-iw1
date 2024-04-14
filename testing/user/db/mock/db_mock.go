// Code generated by MockGen. DO NOT EDIT.
// Source: db.go
//
// Generated by this command:
//
//	mockgen -source=db.go -destination=mock/db_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	model "esgi-iw1/testing/user/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDb is a mock of Db interface.
type MockDb struct {
	ctrl     *gomock.Controller
	recorder *MockDbMockRecorder
}

// MockDbMockRecorder is the mock recorder for MockDb.
type MockDbMockRecorder struct {
	mock *MockDb
}

// NewMockDb creates a new mock instance.
func NewMockDb(ctrl *gomock.Controller) *MockDb {
	mock := &MockDb{ctrl: ctrl}
	mock.recorder = &MockDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDb) EXPECT() *MockDbMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockDb) Exec(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockDbMockRecorder) Exec(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDb)(nil).Exec), arg0, arg1)
}

// Query mocks base method.
func (m *MockDb) Query(arg0 string, arg1 int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDbMockRecorder) Query(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDb)(nil).Query), arg0, arg1)
}
