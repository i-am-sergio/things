package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type MockEnvLoader struct {
    mock.Mock
}
func (m *MockEnvLoader) LoadEnv(filePath string) error {
    args := m.Called(filePath)
    return args.Error(0)
}

type MockDBConnector struct {
    mock.Mock
}
func (m *MockDBConnector) Open(dsn string) (*gorm.DB, error) {
    args := m.Called(dsn)
    return args.Get(0).(*gorm.DB), args.Error(1)
}

func TestGormConnectorOpen(t *testing.T) {
	dsn := "gorm:gorm@tcp(localhost:9910)/gorm?loc=Asia%2FHongKong" 
	connector := GormConnector{}
	_, err := connector.Open(dsn)
	require.Error(t, err, "Expected error connecting to database")
}

func TestDotEnvLoaderLoadEnv(t *testing.T) {
    envContent := []byte("TEST_VAR=success")
	tempEnvFile, err := os.CreateTemp("", "*.env")
	require.NoError(t, err, "Unable to create temporary env file")
	defer os.Remove(tempEnvFile.Name())
	_, err = tempEnvFile.Write(envContent)
	require.NoError(t, err, "Unable to write to temporary env file")
	loader := DotEnvLoader{}
	err = loader.LoadEnv(tempEnvFile.Name())
	require.NoError(t, err, "Error loading .env file")
	val := os.Getenv("TEST_VAR")
	assert.Equal(t, "success", val, "Expected TEST_VAR to be 'success'")
}

func TestGormDBClientInit(t *testing.T) {
	t.Run("SuccessfulInitialization", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockDBConnector := new(MockDBConnector)
		db := new(gorm.DB)
		expectedEnvPath := ".env"
		mockEnvLoader.On("LoadEnv", expectedEnvPath).Return(nil)
		mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, nil)
		err := Init(mockEnvLoader, mockDBConnector)
		require.NoError(t, err)
		mockEnvLoader.AssertExpectations(t)
		mockDBConnector.AssertExpectations(t)
	})

	t.Run("ErrorLoadingEnvFile", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockDBConnector := new(MockDBConnector)
		mockError := fmt.Errorf("failed to load .env file")
		mockEnvLoader.On("LoadEnv",".env").Return(mockError)
		err := Init(mockEnvLoader, mockDBConnector)
		require.Error(t, err)
		expectedError := "failed to initialize database: error loading .env file: failed to load .env file"
		assert.Equal(t, expectedError, err.Error())
		mockEnvLoader.AssertExpectations(t)
		mockDBConnector.AssertNotCalled(t, "Open", mock.AnythingOfType("string"))
	})

	t.Run("ErrorConnectingToDatabase", func(t *testing.T) {
		mockEnvLoader := new(MockEnvLoader)
		mockDBConnector := new(MockDBConnector)
		mockError := fmt.Errorf("failed to connect to database")
		db := new(gorm.DB)
		mockEnvLoader.On("LoadEnv",".env").Return(nil)
		mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, mockError)
		err := Init(mockEnvLoader, mockDBConnector)
		require.Error(t, err)
		expectedError := "failed to initialize database: failed to connect to database: failed to connect to database"
		assert.Equal(t, expectedError, err.Error())
		mockEnvLoader.AssertExpectations(t)
		mockDBConnector.AssertExpectations(t)
	})
}


func TestInit(t *testing.T) {
    t.Run("Success", func(t *testing.T) {
        mockEnvLoader := new(MockEnvLoader)
        mockDBConnector := new(MockDBConnector)
        db := new(gorm.DB)
        mockEnvLoader.On("LoadEnv",".env").Return(nil)
        mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, nil)
        err := Init(mockEnvLoader, mockDBConnector)
        require.NoError(t, err)
        mockEnvLoader.AssertExpectations(t)
        mockDBConnector.AssertExpectations(t)
    })

    t.Run("ErrorInitializingDatabase", func(t *testing.T) {
        mockEnvLoader := new(MockEnvLoader)
        mockDBConnector := new(MockDBConnector)
        mockError := fmt.Errorf("database connection failed")
        mockEnvLoader.On("LoadEnv",".env").Return(nil)
        mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(&gorm.DB{}, mockError) // Devuelve una instancia válida de *gorm.DB
        err := Init(mockEnvLoader, mockDBConnector)
        require.Error(t, err)
        expectedError := fmt.Sprintf("failed to initialize database: failed to connect to database: %v", mockError)
        assert.Equal(t, expectedError, err.Error())
        mockEnvLoader.AssertExpectations(t)
        mockDBConnector.AssertExpectations(t)
    })
}
