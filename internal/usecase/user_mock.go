// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	entitiy "go-mongo-crud-rest-api/internal/entitiy"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSearchIndex is a mock of SearchIndex interface.
type MockSearchIndex struct {
	ctrl     *gomock.Controller
	recorder *MockSearchIndexMockRecorder
}

// MockSearchIndexMockRecorder is the mock recorder for MockSearchIndex.
type MockSearchIndexMockRecorder struct {
	mock *MockSearchIndex
}

// NewMockSearchIndex creates a new mock instance.
func NewMockSearchIndex(ctrl *gomock.Controller) *MockSearchIndex {
	mock := &MockSearchIndex{ctrl: ctrl}
	mock.recorder = &MockSearchIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchIndex) EXPECT() *MockSearchIndexMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockSearchIndex) CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockSearchIndexMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockSearchIndex)(nil).CreateUser), ctx, in)
}

// DeleteUser mocks base method.
func (m *MockSearchIndex) DeleteUser(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockSearchIndexMockRecorder) DeleteUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockSearchIndex)(nil).DeleteUser), ctx, email)
}

// FindUser mocks base method.
func (m *MockSearchIndex) FindUser(ctx context.Context, email string) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, email)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockSearchIndexMockRecorder) FindUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockSearchIndex)(nil).FindUser), ctx, email)
}

// GetUser mocks base method.
func (m *MockSearchIndex) GetUser(ctx context.Context, email string) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockSearchIndexMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockSearchIndex)(nil).GetUser), ctx, email)
}

// UpdateUser mocks base method.
func (m *MockSearchIndex) UpdateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, in)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockSearchIndexMockRecorder) UpdateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockSearchIndex)(nil).UpdateUser), ctx, in)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), ctx, in)
}

// DeleteUser mocks base method.
func (m *MockRepository) DeleteUser(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockRepositoryMockRecorder) DeleteUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockRepository)(nil).DeleteUser), ctx, email)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(ctx context.Context, email string) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), ctx, email)
}

// UpdateUser mocks base method.
func (m *MockRepository) UpdateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, in)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepositoryMockRecorder) UpdateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepository)(nil).UpdateUser), ctx, in)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), ctx, in)
}

// DeleteUser mocks base method.
func (m *MockUser) DeleteUser(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserMockRecorder) DeleteUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUser)(nil).DeleteUser), ctx, email)
}

// FindUser mocks base method.
func (m *MockUser) FindUser(ctx context.Context, email string) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, email)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUserMockRecorder) FindUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUser)(nil).FindUser), ctx, email)
}

// GetUser mocks base method.
func (m *MockUser) GetUser(ctx context.Context, email string) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUser)(nil).GetUser), ctx, email)
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(ctx context.Context, in *entitiy.User) (*entitiy.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, in)
	ret0, _ := ret[0].(*entitiy.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), ctx, in)
}