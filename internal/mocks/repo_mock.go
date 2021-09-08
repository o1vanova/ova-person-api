// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-person-api/internal/repo (interfaces: PersonRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	person "github.com/ozonva/ova-person-api/internal/models"
)

// MockPersonRepo is a mock of PersonRepo interface.
type MockPersonRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPersonRepoMockRecorder
}

// MockPersonRepoMockRecorder is the mock recorder for MockPersonRepo.
type MockPersonRepoMockRecorder struct {
	mock *MockPersonRepo
}

// NewMockPersonRepo creates a new mock instance.
func NewMockPersonRepo(ctrl *gomock.Controller) *MockPersonRepo {
	mock := &MockPersonRepo{ctrl: ctrl}
	mock.recorder = &MockPersonRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonRepo) EXPECT() *MockPersonRepoMockRecorder {
	return m.recorder
}

// AddPerson mocks base method.
func (m *MockPersonRepo) AddPerson(arg0 context.Context, arg1 person.Person) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPerson", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPerson indicates an expected call of AddPerson.
func (mr *MockPersonRepoMockRecorder) AddPerson(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPerson", reflect.TypeOf((*MockPersonRepo)(nil).AddPerson), arg0, arg1)
}

// DescribePerson mocks base method.
func (m *MockPersonRepo) DescribePerson(arg0 context.Context, arg1 uint64) (*person.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribePerson", arg0, arg1)
	ret0, _ := ret[0].(*person.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePerson indicates an expected call of DescribePerson.
func (mr *MockPersonRepoMockRecorder) DescribePerson(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePerson", reflect.TypeOf((*MockPersonRepo)(nil).DescribePerson), arg0, arg1)
}

// ListPersons mocks base method.
func (m *MockPersonRepo) ListPersons(arg0 context.Context, arg1, arg2 uint64) (*[]person.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPersons", arg0, arg1, arg2)
	ret0, _ := ret[0].(*[]person.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPersons indicates an expected call of ListPersons.
func (mr *MockPersonRepoMockRecorder) ListPersons(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPersons", reflect.TypeOf((*MockPersonRepo)(nil).ListPersons), arg0, arg1, arg2)
}

// RemovePerson mocks base method.
func (m *MockPersonRepo) RemovePerson(arg0 context.Context, arg1 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePerson", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemovePerson indicates an expected call of RemovePerson.
func (mr *MockPersonRepoMockRecorder) RemovePerson(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePerson", reflect.TypeOf((*MockPersonRepo)(nil).RemovePerson), arg0, arg1)
}
