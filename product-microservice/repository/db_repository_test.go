package repository

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockGormDB struct {
	mock.Mock
}

func (m *MockGormDB) AutoMigrate(models ...interface{}) error {
	args := m.Called(models)
	return args.Error(0)
}

func (m *MockGormDB) First(dest interface{}, conds ...interface{}) error {
	args := m.Called(dest, []interface{}{conds})
	return args.Error(0)
}

func (m *MockGormDB) Save(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockGormDB) Create(value interface{}) error {
	args := m.Called([]interface{}{value})
	return args.Error(0)
}

func (m *MockGormDB) Preload(relation string) *gorm.DB {
	args := m.Called(relation)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormDB) Find(dest interface{}, conds ...interface{}) error {
	args := m.Called(dest, conds)
	return args.Error(0)
}

func (m *MockGormDB) Where(query string, args ...interface{}) *gorm.DB {
	callArgs := m.Called(query, args)
	return callArgs.Get(0).(*gorm.DB)
}

func (m *MockGormDB) Delete(value interface{}, where ...interface{}) error {
	args := m.Called(value, where)
	return args.Error(0)
}

func (m *MockGormDB) GormDB() *gorm.DB {
	return nil
}
