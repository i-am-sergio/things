package db

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockEnvLoader struct {
    mock.Mock
}
func (m *MockEnvLoader) LoadEnv() error {
    args := m.Called()
    return args.Error(0)
}

type MockDBConnector struct {
    mock.Mock
}
func (m *MockDBConnector) Open(dsn string) (*gorm.DB, error) {
    args := m.Called(dsn)
    return args.Get(0).(*gorm.DB), args.Error(1)
}

func TestGormDBClientInit(t *testing.T) {
	mockEnvLoader := new(MockEnvLoader)
    mockDBConnector := new(MockDBConnector)
    db := new(gorm.DB)
    mockEnvLoader.On("LoadEnv").Return(nil)
    mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, nil)
	err := Init(mockEnvLoader, mockDBConnector)
    if err != nil {
        t.Errorf("Init returned an unexpected error: %v", err)
    }
    mockEnvLoader.AssertExpectations(t)
    mockDBConnector.AssertExpectations(t)
}

func TestInit(t *testing.T) {
    mockEnvLoader := new(MockEnvLoader)
    mockDBConnector := new(MockDBConnector)
    db := new(gorm.DB)
    mockEnvLoader.On("LoadEnv").Return(nil)
    mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, nil)
    Init(mockEnvLoader, mockDBConnector)
    mockEnvLoader.AssertExpectations(t)
    mockDBConnector.AssertExpectations(t)
}