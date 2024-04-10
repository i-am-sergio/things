package db

import (
	"fmt"
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
	t.Run("SuccessfulInitialization", func(t *testing.T) {
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
	})

	t.Run("ErrorLoadingEnvFile", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockDBConnector := new(MockDBConnector)
		mockError := fmt.Errorf("failed to load .env file")
		mockEnvLoader.On("LoadEnv").Return(mockError)
		err := Init(mockEnvLoader, mockDBConnector)
		if err == nil {
			t.Error("Expected an error but got none")
		}
		expectedError := "failed to initialize database: error loading .env file: failed to load .env file"
        if err.Error() != expectedError {
            t.Errorf("Expected error '%s' but got '%v'", expectedError, err)
        }
		mockEnvLoader.AssertExpectations(t)
		mockDBConnector.AssertNotCalled(t, "Open", mock.AnythingOfType("string"))
	})

	t.Run("ErrorConnectingToDatabase", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockDBConnector := new(MockDBConnector)
		mockError := fmt.Errorf("failed to connect to database")
		db := new(gorm.DB)
		mockEnvLoader.On("LoadEnv").Return(nil)
		mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, mockError)
		err := Init(mockEnvLoader, mockDBConnector)
		if err == nil {
			t.Error("Expected an error but got none")
		}
		expectedError := "failed to initialize database: failed to connect to database: failed to connect to database"
        if err.Error() != expectedError {
            t.Errorf("Expected error '%s' but got '%v'", expectedError, err)
        }
		mockEnvLoader.AssertExpectations(t)
		mockDBConnector.AssertExpectations(t)
	})
}


func TestInit(t *testing.T) {
    t.Run("Success", func(t *testing.T) {
        mockEnvLoader := new(MockEnvLoader)
        mockDBConnector := new(MockDBConnector)
        db := new(gorm.DB)
        mockEnvLoader.On("LoadEnv").Return(nil)
        mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, nil)
        err := Init(mockEnvLoader, mockDBConnector)
        if err != nil {
            t.Errorf("Init() error = %v, wantErr %v", err, false)
        }
        mockEnvLoader.AssertExpectations(t)
        mockDBConnector.AssertExpectations(t)
    })

    t.Run("ErrorInitializingDatabase", func(t *testing.T) {
        mockEnvLoader := new(MockEnvLoader)
        mockDBConnector := new(MockDBConnector)
        mockError := fmt.Errorf("database connection failed")
        mockEnvLoader.On("LoadEnv").Return(nil)
        mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(&gorm.DB{}, mockError) // Devuelve una instancia válida de *gorm.DB
        err := Init(mockEnvLoader, mockDBConnector)
        if err == nil {
            t.Fatal("expected error, got nil")
        }
        expectedError := fmt.Sprintf("failed to initialize database: failed to connect to database: %v", mockError)
        if err.Error() != expectedError {
            t.Errorf("expected error %q, got %q", expectedError, err.Error())
        }
        mockEnvLoader.AssertExpectations(t)
        mockDBConnector.AssertExpectations(t)
    })
}
