package db_test

import (
	"auth-microservice/db"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

type MockDBConnector struct {
	mock.Mock
}

func (m *MockDBConnector) DBConnection(dns string) (*gorm.DB, error) {
	args := m.Called(dns)
	return args.Get(0).(*gorm.DB), args.Error(1)

}

func TestDBConnection_Success(t *testing.T) {
	sqlDB, gormDB, mockmock := DbMock(t)
	defer sqlDB.Close()
	mockmock.ExpectPing()

	mockConnector := new(MockDBConnector)
	mockConnector.On("DBConnection", mock.Anything).Return(gormDB, nil)

	dns := "mock_dsn"
	actualDB, err := db.DBConnection(mockConnector, dns)

	assert.NoError(t, err)
	assert.NotNil(t, actualDB)

}

func TestDBConnection_Error(t *testing.T) {
	sqlDB, _, mockmock := DbMock(t)
	defer sqlDB.Close()
	mockmock.ExpectPing()

	mockConnector := new(MockDBConnector)
	mockConnector.On("DBConnection", mock.Anything).Return(&gorm.DB{}, errors.New("Error connecting to db"))

	dns := "mock_dsn"
	actualDB, err := db.DBConnection(mockConnector, dns)

	assert.Error(t, err)
	assert.Nil(t, actualDB)

}
