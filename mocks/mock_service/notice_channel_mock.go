// Code generated by MockGen. DO NOT EDIT.
// Source: service/notice_channel.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	model "github.com/eolinker/apinto-dashboard/model"
	gomock "github.com/golang/mock/gomock"
)

// MockINoticeChannelService is a mock of INoticeChannelService interface.
type MockINoticeChannelService struct {
	ctrl     *gomock.Controller
	recorder *MockINoticeChannelServiceMockRecorder
}

// MockINoticeChannelServiceMockRecorder is the mock recorder for MockINoticeChannelService.
type MockINoticeChannelServiceMockRecorder struct {
	mock *MockINoticeChannelService
}

// NewMockINoticeChannelService creates a new mock instance.
func NewMockINoticeChannelService(ctrl *gomock.Controller) *MockINoticeChannelService {
	mock := &MockINoticeChannelService{ctrl: ctrl}
	mock.recorder = &MockINoticeChannelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINoticeChannelService) EXPECT() *MockINoticeChannelServiceMockRecorder {
	return m.recorder
}

// CreateNoticeChannel mocks base method.
func (m *MockINoticeChannelService) CreateNoticeChannel(ctx context.Context, namespaceId, userID int, channel *model.NoticeChannel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNoticeChannel", ctx, namespaceId, userID, channel)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNoticeChannel indicates an expected call of CreateNoticeChannel.
func (mr *MockINoticeChannelServiceMockRecorder) CreateNoticeChannel(ctx, namespaceId, userID, channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNoticeChannel", reflect.TypeOf((*MockINoticeChannelService)(nil).CreateNoticeChannel), ctx, namespaceId, userID, channel)
}

// DeleteNoticeChannel mocks base method.
func (m *MockINoticeChannelService) DeleteNoticeChannel(ctx context.Context, namespaceId, userID int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNoticeChannel", ctx, namespaceId, userID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNoticeChannel indicates an expected call of DeleteNoticeChannel.
func (mr *MockINoticeChannelServiceMockRecorder) DeleteNoticeChannel(ctx, namespaceId, userID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNoticeChannel", reflect.TypeOf((*MockINoticeChannelService)(nil).DeleteNoticeChannel), ctx, namespaceId, userID, name)
}

// InitChannelDriver mocks base method.
func (m *MockINoticeChannelService) InitChannelDriver() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitChannelDriver")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitChannelDriver indicates an expected call of InitChannelDriver.
func (mr *MockINoticeChannelServiceMockRecorder) InitChannelDriver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitChannelDriver", reflect.TypeOf((*MockINoticeChannelService)(nil).InitChannelDriver))
}

// NoticeChannelByName mocks base method.
func (m *MockINoticeChannelService) NoticeChannelByName(ctx context.Context, namespaceId int, name string) (*model.NoticeChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoticeChannelByName", ctx, namespaceId, name)
	ret0, _ := ret[0].(*model.NoticeChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NoticeChannelByName indicates an expected call of NoticeChannelByName.
func (mr *MockINoticeChannelServiceMockRecorder) NoticeChannelByName(ctx, namespaceId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoticeChannelByName", reflect.TypeOf((*MockINoticeChannelService)(nil).NoticeChannelByName), ctx, namespaceId, name)
}

// NoticeChannelList mocks base method.
func (m *MockINoticeChannelService) NoticeChannelList(ctx context.Context, namespaceId, typ_ int) ([]*model.NoticeChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoticeChannelList", ctx, namespaceId, typ_)
	ret0, _ := ret[0].([]*model.NoticeChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NoticeChannelList indicates an expected call of NoticeChannelList.
func (mr *MockINoticeChannelServiceMockRecorder) NoticeChannelList(ctx, namespaceId, typ_ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoticeChannelList", reflect.TypeOf((*MockINoticeChannelService)(nil).NoticeChannelList), ctx, namespaceId, typ_)
}

// UpdateNoticeChannel mocks base method.
func (m *MockINoticeChannelService) UpdateNoticeChannel(ctx context.Context, namespaceId, userID int, channel *model.NoticeChannel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNoticeChannel", ctx, namespaceId, userID, channel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNoticeChannel indicates an expected call of UpdateNoticeChannel.
func (mr *MockINoticeChannelServiceMockRecorder) UpdateNoticeChannel(ctx, namespaceId, userID, channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNoticeChannel", reflect.TypeOf((*MockINoticeChannelService)(nil).UpdateNoticeChannel), ctx, namespaceId, userID, channel)
}
