// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT
//

// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/communitybridge/easycla/cla-backend-go/gen/v1/models"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddGithubRepository mocks base method
func (m *MockService) AddGithubRepository(ctx context.Context, externalProjectID string, input *models.GithubRepositoryInput) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubAddRepository", ctx, externalProjectID, input)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGithubRepository indicates an expected call of AddGithubRepository
func (mr *MockServiceMockRecorder) AddGithubRepository(ctx, externalProjectID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubAddRepository", reflect.TypeOf((*MockService)(nil).AddGithubRepository), ctx, externalProjectID, input)
}

// EnableRepository mocks base method
func (m *MockService) EnableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubEnableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableRepository indicates an expected call of EnableRepository
func (mr *MockServiceMockRecorder) EnableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubEnableRepository", reflect.TypeOf((*MockService)(nil).EnableRepository), ctx, repositoryID)
}

// EnableRepositoryWithCLAGroupID mocks base method
func (m *MockService) EnableRepositoryWithCLAGroupID(ctx context.Context, repositoryID, claGroupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubEnableRepositoryWithCLAGroupID", ctx, repositoryID, claGroupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableRepositoryWithCLAGroupID indicates an expected call of EnableRepositoryWithCLAGroupID
func (mr *MockServiceMockRecorder) EnableRepositoryWithCLAGroupID(ctx, repositoryID, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubEnableRepositoryWithCLAGroupID", reflect.TypeOf((*MockService)(nil).EnableRepositoryWithCLAGroupID), ctx, repositoryID, claGroupID)
}

// DisableRepository mocks base method
func (m *MockService) DisableRepository(ctx context.Context, repositoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubDisableRepository", ctx, repositoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableRepository indicates an expected call of DisableRepository
func (mr *MockServiceMockRecorder) DisableRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubDisableRepository", reflect.TypeOf((*MockService)(nil).DisableRepository), ctx, repositoryID)
}

// UpdateClaGroupID mocks base method
func (m *MockService) UpdateClaGroupID(ctx context.Context, repositoryID, claGroupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubUpdateClaGroupID", ctx, repositoryID, claGroupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClaGroupID indicates an expected call of UpdateClaGroupID
func (mr *MockServiceMockRecorder) UpdateClaGroupID(ctx, repositoryID, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubUpdateClaGroupID", reflect.TypeOf((*MockService)(nil).UpdateClaGroupID), ctx, repositoryID, claGroupID)
}

// ListProjectRepositories mocks base method
func (m *MockService) ListProjectRepositories(ctx context.Context, externalProjectID string, enabled *bool) (*models.GithubListRepositories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubListProjectRepositories", ctx, externalProjectID, enabled)
	ret0, _ := ret[0].(*models.GithubListRepositories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectRepositories indicates an expected call of ListProjectRepositories
func (mr *MockServiceMockRecorder) ListProjectRepositories(ctx, externalProjectID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubListProjectRepositories", reflect.TypeOf((*MockService)(nil).ListProjectRepositories), ctx, externalProjectID, enabled)
}

// GetRepository mocks base method
func (m *MockService) GetRepository(ctx context.Context, repositoryID string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepository", ctx, repositoryID)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository
func (mr *MockServiceMockRecorder) GetRepository(ctx, repositoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepository", reflect.TypeOf((*MockService)(nil).GetRepository), ctx, repositoryID)
}

// GetRepositoryByProjectSFID mocks base method
func (m *MockService) GetRepositoryByProjectSFID(ctx context.Context, projectSFID string, enabled *bool) (*models.GithubListRepositories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByProjectSFID", ctx, projectSFID, enabled)
	ret0, _ := ret[0].(*models.GithubListRepositories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByProjectSFID indicates an expected call of GetRepositoryByProjectSFID
func (mr *MockServiceMockRecorder) GetRepositoryByProjectSFID(ctx, projectSFID, enabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByProjectSFID", reflect.TypeOf((*MockService)(nil).GetRepositoryByProjectSFID), ctx, projectSFID, enabled)
}

// GetRepositoryByName mocks base method
func (m *MockService) GetRepositoryByName(ctx context.Context, repositoryName string) (*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoryByName", ctx, repositoryName)
	ret0, _ := ret[0].(*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByName indicates an expected call of GetRepositoryByName
func (mr *MockServiceMockRecorder) GetRepositoryByName(ctx, repositoryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepositoryByName", reflect.TypeOf((*MockService)(nil).GetRepositoryByName), ctx, repositoryName)
}

// DisableRepositoriesByProjectID mocks base method
func (m *MockService) DisableRepositoriesByProjectID(ctx context.Context, projectID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubDisableRepositoriesByProjectID", ctx, projectID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableRepositoriesByProjectID indicates an expected call of DisableRepositoriesByProjectID
func (mr *MockServiceMockRecorder) DisableRepositoriesByProjectID(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubDisableRepositoriesByProjectID", reflect.TypeOf((*MockService)(nil).DisableRepositoriesByProjectID), ctx, projectID)
}

// GetRepositoriesByCLAGroup mocks base method
func (m *MockService) GetRepositoriesByCLAGroup(ctx context.Context, claGroupID string) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoriesByCLAGroup", ctx, claGroupID)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByCLAGroup indicates an expected call of GetRepositoriesByCLAGroup
func (mr *MockServiceMockRecorder) GetRepositoriesByCLAGroup(ctx, claGroupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepositoriesByCLAGroup", reflect.TypeOf((*MockService)(nil).GetRepositoriesByCLAGroup), ctx, claGroupID)
}

// GetRepositoriesByOrganizationName mocks base method
func (m *MockService) GetRepositoriesByOrganizationName(ctx context.Context, gitHubOrgName string) ([]*models.GithubRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitHubGetRepositoriesByOrganizationName", ctx, gitHubOrgName)
	ret0, _ := ret[0].([]*models.GithubRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoriesByOrganizationName indicates an expected call of GetRepositoriesByOrganizationName
func (mr *MockServiceMockRecorder) GetRepositoriesByOrganizationName(ctx, gitHubOrgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitHubGetRepositoriesByOrganizationName", reflect.TypeOf((*MockService)(nil).GetRepositoriesByOrganizationName), ctx, gitHubOrgName)
}

// MockGithubOrgRepo is a mock of GithubOrgRepo interface
type MockGithubOrgRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGithubOrgRepoMockRecorder
}

// MockGithubOrgRepoMockRecorder is the mock recorder for MockGithubOrgRepo
type MockGithubOrgRepoMockRecorder struct {
	mock *MockGithubOrgRepo
}

// NewMockGithubOrgRepo creates a new mock instance
func NewMockGithubOrgRepo(ctrl *gomock.Controller) *MockGithubOrgRepo {
	mock := &MockGithubOrgRepo{ctrl: ctrl}
	mock.recorder = &MockGithubOrgRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGithubOrgRepo) EXPECT() *MockGithubOrgRepoMockRecorder {
	return m.recorder
}

// GetGitHubOrganizationByName mocks base method
func (m *MockGithubOrgRepo) GetGitHubOrganizationByName(ctx context.Context, githubOrganizationName string) (*models.GithubOrganizations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitHubOrganizationByName", ctx, githubOrganizationName)
	ret0, _ := ret[0].(*models.GithubOrganizations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGithubOrganizationByName indicates an expected call of GetGithubOrganizationByName
func (mr *MockGithubOrgRepoMockRecorder) GetGithubOrganizationByName(ctx, githubOrganizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitHubOrganizationByName", reflect.TypeOf((*MockGithubOrgRepo)(nil).GetGitHubOrganizationByName), ctx, githubOrganizationName)
}

// GetGithubOrganization mocks base method
func (m *MockGithubOrgRepo) GetGitHubOrganization(ctx context.Context, githubOrganizationName string) (*models.GithubOrganization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitHubOrganization", ctx, githubOrganizationName)
	ret0, _ := ret[0].(*models.GithubOrganization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGithubOrganization indicates an expected call of GetGithubOrganization
func (mr *MockGithubOrgRepoMockRecorder) GetGithubOrganization(ctx, githubOrganizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitHubOrganization", reflect.TypeOf((*MockGithubOrgRepo)(nil).GetGitHubOrganization), ctx, githubOrganizationName)
}

// GetGitHubOrganizations mocks base method
func (m *MockGithubOrgRepo) GetGitHubOrganizations(ctx context.Context, projectSFID string) (*models.GithubOrganizations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitHubOrganizations", ctx, projectSFID)
	ret0, _ := ret[0].(*models.GithubOrganizations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGithubOrganizations indicates an expected call of GetGithubOrganizations
func (mr *MockGithubOrgRepoMockRecorder) GetGithubOrganizations(ctx, projectSFID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitHubOrganizations", reflect.TypeOf((*MockGithubOrgRepo)(nil).GetGitHubOrganizations), ctx, projectSFID)
}
