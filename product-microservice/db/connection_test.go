package db

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)
type MockDBInterface struct {
    mock.Mock
}

type MockEnvLoader struct {
    mock.Mock
}

type MockDBConnector struct {
    mock.Mock
}

func (m *MockDBConnector) Open(dsn string) (*gorm.DB, error) {
    args := m.Called(dsn)
    return args.Get(0).(*gorm.DB), args.Error(1)
}

func (m *MockDBInterface) AutoMigrate(dst ...interface{}) error {
    args := m.Called(dst)
    return args.Error(0)
}

func (m *MockDBInterface) First(dest interface{}, conds ...interface{}) error {
    args := m.Called(dest, conds)
    return args.Error(0)
}

func (m *MockDBInterface) Save(value interface{}) error {
    args := m.Called(value)
    return args.Error(0)
}

func (m *MockDBInterface) Create(value interface{}) error {
    args := m.Called(value)
    return args.Error(0)
}

func (m *MockDBInterface) FindPreloaded(relation string, dest interface{}, conds ...interface{}) error {
    args := m.Called(relation, dest, conds)
    return args.Error(0)
}

func (m *MockDBInterface) Find(dest interface{}, conds ...interface{}) error {
    args := m.Called(dest, conds)
    return args.Error(0)
}

func (m *MockDBInterface) FindWithCondition(dest interface{}, query string, args ...interface{}) error {
    mockArgs := m.Called(dest, query, args)
    return mockArgs.Error(0)
}

func (m *MockDBInterface) Delete(value interface{}) error {
    args := m.Called(value)
    return args.Error(0)
}

func (m *MockDBInterface) DeleteWithCondition(model interface{}, query string, args ...interface{}) error {
    mockArgs := m.Called(model, query, args)
    return mockArgs.Error(0)
}

func (m *MockDBInterface) DeleteByID(model interface{}, id interface{}) error {
    mockArgs := m.Called(model, id)
    return mockArgs.Error(0)
}

func (m *MockEnvLoader) LoadEnv() error {
    args := m.Called()
    return args.Error(0)
}

func TestGormDBClientInit(t *testing.T) {
	mockEnvLoader := new(MockEnvLoader)
    mockDBConnector := new(MockDBConnector)
    db := new(gorm.DB)
    mockEnvLoader.On("LoadEnv").Return(nil)
    mockDBConnector.On("Open", mock.AnythingOfType("string")).Return(db, nil)
	gormDBClient := GormDBClient{EnvLoader: mockEnvLoader, Connector: mockDBConnector}
    gormDBClient.Init(mockEnvLoader)
    mockEnvLoader.AssertExpectations(t)
    mockDBConnector.AssertExpectations(t)
}
