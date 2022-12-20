// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/lawyer/model/lawyer.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	reflect "reflect"

	err "github.com/dougmendes/advg/pkg/errs"
	model "github.com/dougmendes/advg/pkg/lawyer/model"
	gomock "github.com/golang/mock/gomock"
)

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

// GetById mocks base method.
func (m *MockRepository) GetById(ctx context.Context, id int) (*model.Lawyer, *err.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*model.Lawyer)
	ret1, _ := ret[1].(*err.Error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRepositoryMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), ctx, id)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetLawyerById mocks base method.
func (m *MockService) GetLawyerById(ctx context.Context, id int) (*model.Lawyer, *err.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLawyerById", ctx, id)
	ret0, _ := ret[0].(*model.Lawyer)
	ret1, _ := ret[1].(*err.Error)
	return ret0, ret1
}

// GetLawyerById indicates an expected call of GetLawyerById.
func (mr *MockServiceMockRecorder) GetLawyerById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLawyerById", reflect.TypeOf((*MockService)(nil).GetLawyerById), ctx, id)
}
