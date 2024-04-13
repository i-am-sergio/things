package db_test

import (
	"auth-microservice/db"
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	testDNS := "host=monorail.proxy.rlwy.net user=postgres password=eWhJpljFwQgFGkhMTTpOVfFCpdqhWLMY dbname=railway port=24696"
	dbInstance, err := db.DBConnection(testDNS)
	if err != nil {
		log.Println("Error connecting to the database:", err)
	}
	assert.NotNil(t, dbInstance, "Database connection should not be nil")
	assert.NoError(t, err, "Failed to connect to the database")
	if dbInstance != nil {
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		assert.NoError(t, err, "Failed to close the database connection")

	}

}

func TestConnectDBError(t *testing.T) {
	testDNS := "host=localhostt user=postgres password=admin dbname=users port=5432"
	instanceDB, err := db.DBConnection(testDNS)
	if err != nil {
		log.Println("Error connecting to the database:", err)
	}
	assert.Error(t, err)
	assert.Nil(t, instanceDB)

}

// func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
// 	sqldb, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	gormdb, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: sqldb,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return sqldb, gormdb, mock
// }

// type MockDBConnector struct {
// 	mock.Mock
// }

// func (m *MockDBConnector) DBConnection(dns string) (*gorm.DB, error) {
// 	args := m.Called(dns)
// 	return args.Get(0).(*gorm.DB), args.Error(1)

// }

// func TestDBConnection_Success(t *testing.T) {
// 	sqlDB, gormDB, mockmock := DbMock(t)
// 	defer sqlDB.Close()
// 	mockmock.ExpectPing()

// 	mockConnector := new(MockDBConnector)
// 	mockConnector.On("DBConnection", mock.Anything).Return(gormDB, nil)

// 	dns := "mock_dsn"
// 	actualDB, err := db.DBConnection(mockConnector, dns)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, actualDB)

// }

// func TestDBConnection_Error(t *testing.T) {
// 	sqlDB, _, mockmock := DbMock(t)
// 	defer sqlDB.Close()
// 	mockmock.ExpectPing()

// 	mockConnector := new(MockDBConnector)
// 	mockConnector.On("DBConnection", mock.Anything).Return(&gorm.DB{}, errors.New("Error connecting to db"))

// 	dns := "mock_dsn"
// 	actualDB, err := db.DBConnection(mockConnector, dns)

// 	assert.Error(t, err)
// 	assert.Nil(t, actualDB)

// }
