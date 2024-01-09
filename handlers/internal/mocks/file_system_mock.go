// Code generated by MockGen. DO NOT EDIT.
// Source: file_system.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=../mocks/file_system_mock.go -source=file_system.go -mock_names=Filesystem=MockFilesystem
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	fsnotify "github.com/fsnotify/fsnotify"
	filesystem "github.com/k-lb/entrypoint-framework/handlers/internal/filesystem"
	gomock "go.uber.org/mock/gomock"
)

// MockFilesystem is a mock of Filesystem interface.
type MockFilesystem struct {
	ctrl     *gomock.Controller
	recorder *MockFilesystemMockRecorder
}

// MockFilesystemMockRecorder is the mock recorder for MockFilesystem.
type MockFilesystemMockRecorder struct {
	mock *MockFilesystem
}

// NewMockFilesystem creates a new mock instance.
func NewMockFilesystem(ctrl *gomock.Controller) *MockFilesystem {
	mock := &MockFilesystem{ctrl: ctrl}
	mock.recorder = &MockFilesystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilesystem) EXPECT() *MockFilesystemMockRecorder {
	return m.recorder
}

// AreFilesDifferent mocks base method.
func (m *MockFilesystem) AreFilesDifferent(firstFilePath, secondFilePath string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreFilesDifferent", firstFilePath, secondFilePath)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AreFilesDifferent indicates an expected call of AreFilesDifferent.
func (mr *MockFilesystemMockRecorder) AreFilesDifferent(firstFilePath, secondFilePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreFilesDifferent", reflect.TypeOf((*MockFilesystem)(nil).AreFilesDifferent), firstFilePath, secondFilePath)
}

// ClearDir mocks base method.
func (m *MockFilesystem) ClearDir(filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearDir", filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearDir indicates an expected call of ClearDir.
func (mr *MockFilesystemMockRecorder) ClearDir(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearDir", reflect.TypeOf((*MockFilesystem)(nil).ClearDir), filePath)
}

// Copy mocks base method.
func (m *MockFilesystem) Copy(fromPath, toPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", fromPath, toPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockFilesystemMockRecorder) Copy(fromPath, toPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockFilesystem)(nil).Copy), fromPath, toPath)
}

// DeleteFile mocks base method.
func (m *MockFilesystem) DeleteFile(filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockFilesystemMockRecorder) DeleteFile(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFilesystem)(nil).DeleteFile), filePath)
}

// DoesExist mocks base method.
func (m *MockFilesystem) DoesExist(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesExist", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DoesExist indicates an expected call of DoesExist.
func (mr *MockFilesystemMockRecorder) DoesExist(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesExist", reflect.TypeOf((*MockFilesystem)(nil).DoesExist), path)
}

// Extract mocks base method.
func (m *MockFilesystem) Extract(tarball, toDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Extract", tarball, toDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// Extract indicates an expected call of Extract.
func (mr *MockFilesystemMockRecorder) Extract(tarball, toDir any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Extract", reflect.TypeOf((*MockFilesystem)(nil).Extract), tarball, toDir)
}

// Hardlink mocks base method.
func (m *MockFilesystem) Hardlink(filePath, hardlinkPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hardlink", filePath, hardlinkPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Hardlink indicates an expected call of Hardlink.
func (mr *MockFilesystemMockRecorder) Hardlink(filePath, hardlinkPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hardlink", reflect.TypeOf((*MockFilesystem)(nil).Hardlink), filePath, hardlinkPath)
}

// ListFileNamesInDir mocks base method.
func (m *MockFilesystem) ListFileNamesInDir(dirPath string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFileNamesInDir", dirPath)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFileNamesInDir indicates an expected call of ListFileNamesInDir.
func (mr *MockFilesystemMockRecorder) ListFileNamesInDir(dirPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFileNamesInDir", reflect.TypeOf((*MockFilesystem)(nil).ListFileNamesInDir), dirPath)
}

// MoveFile mocks base method.
func (m *MockFilesystem) MoveFile(fromPath, toPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveFile", fromPath, toPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveFile indicates an expected call of MoveFile.
func (mr *MockFilesystemMockRecorder) MoveFile(fromPath, toPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveFile", reflect.TypeOf((*MockFilesystem)(nil).MoveFile), fromPath, toPath)
}

// NewFileWatcher mocks base method.
func (m *MockFilesystem) NewFileWatcher(watchedFile string, watchedOps fsnotify.Op) (filesystem.Watcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFileWatcher", watchedFile, watchedOps)
	ret0, _ := ret[0].(filesystem.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFileWatcher indicates an expected call of NewFileWatcher.
func (mr *MockFilesystemMockRecorder) NewFileWatcher(watchedFile, watchedOps any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFileWatcher", reflect.TypeOf((*MockFilesystem)(nil).NewFileWatcher), watchedFile, watchedOps)
}
