// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/focalboard/server/services/store (interfaces: Store)

// Package mockstore is a generated GoMock package.
package mockstore

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/focalboard/server/model"
	store "github.com/mattermost/focalboard/server/services/store"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CleanUpSessions mocks base method.
func (m *MockStore) CleanUpSessions(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUpSessions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUpSessions indicates an expected call of CleanUpSessions.
func (mr *MockStoreMockRecorder) CleanUpSessions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUpSessions", reflect.TypeOf((*MockStore)(nil).CleanUpSessions), arg0)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0)
}

// CreateSubscription mocks base method.
func (m *MockStore) CreateSubscription(arg0 store.Container, arg1 *model.Subscription) (*model.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", arg0, arg1)
	ret0, _ := ret[0].(*model.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockStoreMockRecorder) CreateSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockStore)(nil).CreateSubscription), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0)
}

// DBType mocks base method.
func (m *MockStore) DBType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBType")
	ret0, _ := ret[0].(string)
	return ret0
}

// DBType indicates an expected call of DBType.
func (mr *MockStoreMockRecorder) DBType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBType", reflect.TypeOf((*MockStore)(nil).DBType))
}

// DeleteBlock mocks base method.
func (m *MockStore) DeleteBlock(arg0 store.Container, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlock indicates an expected call of DeleteBlock.
func (mr *MockStoreMockRecorder) DeleteBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlock", reflect.TypeOf((*MockStore)(nil).DeleteBlock), arg0, arg1, arg2)
}

// DeleteNotificationHint mocks base method.
func (m *MockStore) DeleteNotificationHint(arg0 store.Container, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNotificationHint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNotificationHint indicates an expected call of DeleteNotificationHint.
func (mr *MockStoreMockRecorder) DeleteNotificationHint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNotificationHint", reflect.TypeOf((*MockStore)(nil).DeleteNotificationHint), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0)
}

// DeleteSubscription mocks base method.
func (m *MockStore) DeleteSubscription(arg0 store.Container, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription.
func (mr *MockStoreMockRecorder) DeleteSubscription(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockStore)(nil).DeleteSubscription), arg0, arg1, arg2)
}

// GetActiveUserCount mocks base method.
func (m *MockStore) GetActiveUserCount(arg0 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveUserCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveUserCount indicates an expected call of GetActiveUserCount.
func (mr *MockStoreMockRecorder) GetActiveUserCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveUserCount", reflect.TypeOf((*MockStore)(nil).GetActiveUserCount), arg0)
}

// GetAllBlocks mocks base method.
func (m *MockStore) GetAllBlocks(arg0 store.Container) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBlocks", arg0)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBlocks indicates an expected call of GetAllBlocks.
func (mr *MockStoreMockRecorder) GetAllBlocks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBlocks", reflect.TypeOf((*MockStore)(nil).GetAllBlocks), arg0)
}

// GetBlock mocks base method.
func (m *MockStore) GetBlock(arg0 store.Container, arg1 string) (*model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockStoreMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockStore)(nil).GetBlock), arg0, arg1)
}

// GetBlockCountsByType mocks base method.
func (m *MockStore) GetBlockCountsByType() (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockCountsByType")
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockCountsByType indicates an expected call of GetBlockCountsByType.
func (mr *MockStoreMockRecorder) GetBlockCountsByType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockCountsByType", reflect.TypeOf((*MockStore)(nil).GetBlockCountsByType))
}

// GetBlockHistory mocks base method.
func (m *MockStore) GetBlockHistory(arg0 store.Container, arg1 string, arg2 model.QueryBlockHistoryOptions) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHistory", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockHistory indicates an expected call of GetBlockHistory.
func (mr *MockStoreMockRecorder) GetBlockHistory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHistory", reflect.TypeOf((*MockStore)(nil).GetBlockHistory), arg0, arg1, arg2)
}

// GetBlocksWithParent mocks base method.
func (m *MockStore) GetBlocksWithParent(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithParent", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithParent indicates an expected call of GetBlocksWithParent.
func (mr *MockStoreMockRecorder) GetBlocksWithParent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithParent", reflect.TypeOf((*MockStore)(nil).GetBlocksWithParent), arg0, arg1)
}

// GetBlocksWithParentAndType mocks base method.
func (m *MockStore) GetBlocksWithParentAndType(arg0 store.Container, arg1, arg2 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithParentAndType", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithParentAndType indicates an expected call of GetBlocksWithParentAndType.
func (mr *MockStoreMockRecorder) GetBlocksWithParentAndType(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithParentAndType", reflect.TypeOf((*MockStore)(nil).GetBlocksWithParentAndType), arg0, arg1, arg2)
}

// GetBlocksWithRootID mocks base method.
func (m *MockStore) GetBlocksWithRootID(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithRootID", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithRootID indicates an expected call of GetBlocksWithRootID.
func (mr *MockStoreMockRecorder) GetBlocksWithRootID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithRootID", reflect.TypeOf((*MockStore)(nil).GetBlocksWithRootID), arg0, arg1)
}

// GetBlocksWithType mocks base method.
func (m *MockStore) GetBlocksWithType(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithType", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithType indicates an expected call of GetBlocksWithType.
func (mr *MockStoreMockRecorder) GetBlocksWithType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithType", reflect.TypeOf((*MockStore)(nil).GetBlocksWithType), arg0, arg1)
}

// GetBoardAndCard mocks base method.
func (m *MockStore) GetBoardAndCard(arg0 store.Container, arg1 *model.Block) (*model.Block, *model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardAndCard", arg0, arg1)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(*model.Block)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBoardAndCard indicates an expected call of GetBoardAndCard.
func (mr *MockStoreMockRecorder) GetBoardAndCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardAndCard", reflect.TypeOf((*MockStore)(nil).GetBoardAndCard), arg0, arg1)
}

// GetBoardAndCardByID mocks base method.
func (m *MockStore) GetBoardAndCardByID(arg0 store.Container, arg1 string) (*model.Block, *model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardAndCardByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(*model.Block)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBoardAndCardByID indicates an expected call of GetBoardAndCardByID.
func (mr *MockStoreMockRecorder) GetBoardAndCardByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardAndCardByID", reflect.TypeOf((*MockStore)(nil).GetBoardAndCardByID), arg0, arg1)
}

// GetDefaultTemplateBlocks mocks base method.
func (m *MockStore) GetDefaultTemplateBlocks() ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultTemplateBlocks")
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultTemplateBlocks indicates an expected call of GetDefaultTemplateBlocks.
func (mr *MockStoreMockRecorder) GetDefaultTemplateBlocks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultTemplateBlocks", reflect.TypeOf((*MockStore)(nil).GetDefaultTemplateBlocks))
}

// GetNextNotificationHint mocks base method.
func (m *MockStore) GetNextNotificationHint(arg0 bool) (*model.NotificationHint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextNotificationHint", arg0)
	ret0, _ := ret[0].(*model.NotificationHint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextNotificationHint indicates an expected call of GetNextNotificationHint.
func (mr *MockStoreMockRecorder) GetNextNotificationHint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextNotificationHint", reflect.TypeOf((*MockStore)(nil).GetNextNotificationHint), arg0)
}

// GetNotificationHint mocks base method.
func (m *MockStore) GetNotificationHint(arg0 store.Container, arg1 string) (*model.NotificationHint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotificationHint", arg0, arg1)
	ret0, _ := ret[0].(*model.NotificationHint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotificationHint indicates an expected call of GetNotificationHint.
func (mr *MockStoreMockRecorder) GetNotificationHint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotificationHint", reflect.TypeOf((*MockStore)(nil).GetNotificationHint), arg0, arg1)
}

// GetParentID mocks base method.
func (m *MockStore) GetParentID(arg0 store.Container, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParentID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParentID indicates an expected call of GetParentID.
func (mr *MockStoreMockRecorder) GetParentID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentID", reflect.TypeOf((*MockStore)(nil).GetParentID), arg0, arg1)
}

// GetRegisteredUserCount mocks base method.
func (m *MockStore) GetRegisteredUserCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegisteredUserCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegisteredUserCount indicates an expected call of GetRegisteredUserCount.
func (mr *MockStoreMockRecorder) GetRegisteredUserCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegisteredUserCount", reflect.TypeOf((*MockStore)(nil).GetRegisteredUserCount))
}

// GetRootID mocks base method.
func (m *MockStore) GetRootID(arg0 store.Container, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRootID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRootID indicates an expected call of GetRootID.
func (mr *MockStoreMockRecorder) GetRootID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRootID", reflect.TypeOf((*MockStore)(nil).GetRootID), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 string, arg1 int64) (*model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetSharing mocks base method.
func (m *MockStore) GetSharing(arg0 store.Container, arg1 string) (*model.Sharing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharing", arg0, arg1)
	ret0, _ := ret[0].(*model.Sharing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharing indicates an expected call of GetSharing.
func (mr *MockStoreMockRecorder) GetSharing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharing", reflect.TypeOf((*MockStore)(nil).GetSharing), arg0, arg1)
}

// GetSubTree2 mocks base method.
func (m *MockStore) GetSubTree2(arg0 store.Container, arg1 string, arg2 model.QuerySubtreeOptions) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubTree2", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubTree2 indicates an expected call of GetSubTree2.
func (mr *MockStoreMockRecorder) GetSubTree2(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubTree2", reflect.TypeOf((*MockStore)(nil).GetSubTree2), arg0, arg1, arg2)
}

// GetSubTree3 mocks base method.
func (m *MockStore) GetSubTree3(arg0 store.Container, arg1 string, arg2 model.QuerySubtreeOptions) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubTree3", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubTree3 indicates an expected call of GetSubTree3.
func (mr *MockStoreMockRecorder) GetSubTree3(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubTree3", reflect.TypeOf((*MockStore)(nil).GetSubTree3), arg0, arg1, arg2)
}

// GetSubscribersCountForBlock mocks base method.
func (m *MockStore) GetSubscribersCountForBlock(arg0 store.Container, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribersCountForBlock", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribersCountForBlock indicates an expected call of GetSubscribersCountForBlock.
func (mr *MockStoreMockRecorder) GetSubscribersCountForBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribersCountForBlock", reflect.TypeOf((*MockStore)(nil).GetSubscribersCountForBlock), arg0, arg1)
}

// GetSubscribersForBlock mocks base method.
func (m *MockStore) GetSubscribersForBlock(arg0 store.Container, arg1 string) ([]*model.Subscriber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribersForBlock", arg0, arg1)
	ret0, _ := ret[0].([]*model.Subscriber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribersForBlock indicates an expected call of GetSubscribersForBlock.
func (mr *MockStoreMockRecorder) GetSubscribersForBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribersForBlock", reflect.TypeOf((*MockStore)(nil).GetSubscribersForBlock), arg0, arg1)
}

// GetSubscription mocks base method.
func (m *MockStore) GetSubscription(arg0 store.Container, arg1, arg2 string) (*model.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscription", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscription indicates an expected call of GetSubscription.
func (mr *MockStoreMockRecorder) GetSubscription(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscription", reflect.TypeOf((*MockStore)(nil).GetSubscription), arg0, arg1, arg2)
}

// GetSubscriptions mocks base method.
func (m *MockStore) GetSubscriptions(arg0 store.Container, arg1 string) ([]*model.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptions", arg0, arg1)
	ret0, _ := ret[0].([]*model.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions.
func (mr *MockStoreMockRecorder) GetSubscriptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptions", reflect.TypeOf((*MockStore)(nil).GetSubscriptions), arg0, arg1)
}

// GetSystemSetting mocks base method.
func (m *MockStore) GetSystemSetting(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemSetting", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemSetting indicates an expected call of GetSystemSetting.
func (mr *MockStoreMockRecorder) GetSystemSetting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemSetting", reflect.TypeOf((*MockStore)(nil).GetSystemSetting), arg0)
}

// GetSystemSettings mocks base method.
func (m *MockStore) GetSystemSettings() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemSettings")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemSettings indicates an expected call of GetSystemSettings.
func (mr *MockStoreMockRecorder) GetSystemSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemSettings", reflect.TypeOf((*MockStore)(nil).GetSystemSettings))
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), arg0)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0)
}

// GetUserWorkspaces mocks base method.
func (m *MockStore) GetUserWorkspaces(arg0 string) ([]model.UserWorkspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWorkspaces", arg0)
	ret0, _ := ret[0].([]model.UserWorkspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWorkspaces indicates an expected call of GetUserWorkspaces.
func (mr *MockStoreMockRecorder) GetUserWorkspaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWorkspaces", reflect.TypeOf((*MockStore)(nil).GetUserWorkspaces), arg0)
}

// GetUsersByWorkspace mocks base method.
func (m *MockStore) GetUsersByWorkspace(arg0 string) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByWorkspace", arg0)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByWorkspace indicates an expected call of GetUsersByWorkspace.
func (mr *MockStoreMockRecorder) GetUsersByWorkspace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByWorkspace", reflect.TypeOf((*MockStore)(nil).GetUsersByWorkspace), arg0)
}

// GetWorkspace mocks base method.
func (m *MockStore) GetWorkspace(arg0 string) (*model.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", arg0)
	ret0, _ := ret[0].(*model.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockStoreMockRecorder) GetWorkspace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockStore)(nil).GetWorkspace), arg0)
}

// GetWorkspaceCount mocks base method.
func (m *MockStore) GetWorkspaceCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaceCount indicates an expected call of GetWorkspaceCount.
func (mr *MockStoreMockRecorder) GetWorkspaceCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceCount", reflect.TypeOf((*MockStore)(nil).GetWorkspaceCount))
}

// HasWorkspaceAccess mocks base method.
func (m *MockStore) HasWorkspaceAccess(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasWorkspaceAccess", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasWorkspaceAccess indicates an expected call of HasWorkspaceAccess.
func (mr *MockStoreMockRecorder) HasWorkspaceAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasWorkspaceAccess", reflect.TypeOf((*MockStore)(nil).HasWorkspaceAccess), arg0, arg1)
}

// InsertBlock mocks base method.
func (m *MockStore) InsertBlock(arg0 store.Container, arg1 *model.Block, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBlock indicates an expected call of InsertBlock.
func (mr *MockStoreMockRecorder) InsertBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlock", reflect.TypeOf((*MockStore)(nil).InsertBlock), arg0, arg1, arg2)
}

// InsertBlocks mocks base method.
func (m *MockStore) InsertBlocks(arg0 store.Container, arg1 []model.Block, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlocks", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBlocks indicates an expected call of InsertBlocks.
func (mr *MockStoreMockRecorder) InsertBlocks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlocks", reflect.TypeOf((*MockStore)(nil).InsertBlocks), arg0, arg1, arg2)
}

// IsErrNotFound mocks base method.
func (m *MockStore) IsErrNotFound(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsErrNotFound", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsErrNotFound indicates an expected call of IsErrNotFound.
func (mr *MockStoreMockRecorder) IsErrNotFound(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsErrNotFound", reflect.TypeOf((*MockStore)(nil).IsErrNotFound), arg0)
}

// PatchBlock mocks base method.
func (m *MockStore) PatchBlock(arg0 store.Container, arg1 string, arg2 *model.BlockPatch, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchBlock", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchBlock indicates an expected call of PatchBlock.
func (mr *MockStoreMockRecorder) PatchBlock(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBlock", reflect.TypeOf((*MockStore)(nil).PatchBlock), arg0, arg1, arg2, arg3)
}

// PatchBlocks mocks base method.
func (m *MockStore) PatchBlocks(arg0 store.Container, arg1 *model.BlockPatchBatch, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchBlocks", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchBlocks indicates an expected call of PatchBlocks.
func (mr *MockStoreMockRecorder) PatchBlocks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBlocks", reflect.TypeOf((*MockStore)(nil).PatchBlocks), arg0, arg1, arg2)
}

// RefreshSession mocks base method.
func (m *MockStore) RefreshSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshSession indicates an expected call of RefreshSession.
func (mr *MockStoreMockRecorder) RefreshSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshSession", reflect.TypeOf((*MockStore)(nil).RefreshSession), arg0)
}

// RemoveDefaultTemplates mocks base method.
func (m *MockStore) RemoveDefaultTemplates(arg0 []model.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDefaultTemplates", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDefaultTemplates indicates an expected call of RemoveDefaultTemplates.
func (mr *MockStoreMockRecorder) RemoveDefaultTemplates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDefaultTemplates", reflect.TypeOf((*MockStore)(nil).RemoveDefaultTemplates), arg0)
}

// SetSystemSetting mocks base method.
func (m *MockStore) SetSystemSetting(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSystemSetting", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSystemSetting indicates an expected call of SetSystemSetting.
func (mr *MockStoreMockRecorder) SetSystemSetting(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSystemSetting", reflect.TypeOf((*MockStore)(nil).SetSystemSetting), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockStore) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockStoreMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockStore)(nil).Shutdown))
}

// UndeleteBlock mocks base method.
func (m *MockStore) UndeleteBlock(arg0 store.Container, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndeleteBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UndeleteBlock indicates an expected call of UndeleteBlock.
func (mr *MockStoreMockRecorder) UndeleteBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndeleteBlock", reflect.TypeOf((*MockStore)(nil).UndeleteBlock), arg0, arg1, arg2)
}

// UpdateSession mocks base method.
func (m *MockStore) UpdateSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockStoreMockRecorder) UpdateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockStore)(nil).UpdateSession), arg0)
}

// UpdateSubscribersNotifiedAt mocks base method.
func (m *MockStore) UpdateSubscribersNotifiedAt(arg0 store.Container, arg1 string, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscribersNotifiedAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscribersNotifiedAt indicates an expected call of UpdateSubscribersNotifiedAt.
func (mr *MockStoreMockRecorder) UpdateSubscribersNotifiedAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscribersNotifiedAt", reflect.TypeOf((*MockStore)(nil).UpdateSubscribersNotifiedAt), arg0, arg1, arg2)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), arg0, arg1)
}

// UpdateUserPasswordByID mocks base method.
func (m *MockStore) UpdateUserPasswordByID(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPasswordByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPasswordByID indicates an expected call of UpdateUserPasswordByID.
func (mr *MockStoreMockRecorder) UpdateUserPasswordByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPasswordByID", reflect.TypeOf((*MockStore)(nil).UpdateUserPasswordByID), arg0, arg1)
}

// UpsertNotificationHint mocks base method.
func (m *MockStore) UpsertNotificationHint(arg0 *model.NotificationHint, arg1 time.Duration) (*model.NotificationHint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNotificationHint", arg0, arg1)
	ret0, _ := ret[0].(*model.NotificationHint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertNotificationHint indicates an expected call of UpsertNotificationHint.
func (mr *MockStoreMockRecorder) UpsertNotificationHint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNotificationHint", reflect.TypeOf((*MockStore)(nil).UpsertNotificationHint), arg0, arg1)
}

// UpsertSharing mocks base method.
func (m *MockStore) UpsertSharing(arg0 store.Container, arg1 model.Sharing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSharing", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSharing indicates an expected call of UpsertSharing.
func (mr *MockStoreMockRecorder) UpsertSharing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSharing", reflect.TypeOf((*MockStore)(nil).UpsertSharing), arg0, arg1)
}

// UpsertWorkspaceSettings mocks base method.
func (m *MockStore) UpsertWorkspaceSettings(arg0 model.Workspace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkspaceSettings", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkspaceSettings indicates an expected call of UpsertWorkspaceSettings.
func (mr *MockStoreMockRecorder) UpsertWorkspaceSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkspaceSettings", reflect.TypeOf((*MockStore)(nil).UpsertWorkspaceSettings), arg0)
}

// UpsertWorkspaceSignupToken mocks base method.
func (m *MockStore) UpsertWorkspaceSignupToken(arg0 model.Workspace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkspaceSignupToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkspaceSignupToken indicates an expected call of UpsertWorkspaceSignupToken.
func (mr *MockStoreMockRecorder) UpsertWorkspaceSignupToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkspaceSignupToken", reflect.TypeOf((*MockStore)(nil).UpsertWorkspaceSignupToken), arg0)
}
