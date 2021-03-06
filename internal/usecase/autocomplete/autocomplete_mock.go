// Code generated by MockGen. DO NOT EDIT.
// Source: ./autocomplete.go

// Package autocomplete is a generated GoMock package.
package autocomplete

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	search "github.com/ryeoman/go_dependency_injection/internal/repository/search"
)

// MockSearchProvider is a mock of SearchProvider interface.
type MockSearchProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSearchProviderMockRecorder
}

// MockSearchProviderMockRecorder is the mock recorder for MockSearchProvider.
type MockSearchProviderMockRecorder struct {
	mock *MockSearchProvider
}

// NewMockSearchProvider creates a new mock instance.
func NewMockSearchProvider(ctrl *gomock.Controller) *MockSearchProvider {
	mock := &MockSearchProvider{ctrl: ctrl}
	mock.recorder = &MockSearchProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchProvider) EXPECT() *MockSearchProviderMockRecorder {
	return m.recorder
}

// Suggestion mocks base method.
func (m *MockSearchProvider) Suggestion(ctx context.Context, keyword string) (search.Words, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suggestion", ctx, keyword)
	ret0, _ := ret[0].(search.Words)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Suggestion indicates an expected call of Suggestion.
func (mr *MockSearchProviderMockRecorder) Suggestion(ctx, keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suggestion", reflect.TypeOf((*MockSearchProvider)(nil).Suggestion), ctx, keyword)
}
