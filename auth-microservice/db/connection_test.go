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
	// Preparar el mock de la base de datos y el mock de GORM
	sqlDB, gormDB, mockmock := DbMock(t)
	defer sqlDB.Close() // Asegúrate de cerrar la conexión a la base de datos simulada

	// Configurar las expectativas del mock
	mockmock.ExpectPing()

	// Crear un mock de DBConnector
	mockConnector := new(MockDBConnector)
	mockConnector.On("DBConnection", mock.Anything).Return(gormDB, nil)

	// Intentar conectar a la base de datos utilizando el mock de DBConnector
	dns := "mock_dsn" // Esto podría ser cualquier cadena ya que estamos usando un mock
	actualDB, err := db.DBConnection(mockConnector, dns)

	// Verificar que no hay errores y que la conexión se haya establecido correctamente
	assert.NoError(t, err)
	assert.NotNil(t, actualDB)

	// Verificar que todas las expectativas del mock se cumplieron
}

func TestDBConnection_Error(t *testing.T) {
	// Preparar el mock de la base de datos y el mock de GORM
	sqlDB, _, mockmock := DbMock(t)
	defer sqlDB.Close() // Asegúrate de cerrar la conexión a la base de datos simulada

	// Configurar las expectativas del mock
	mockmock.ExpectPing()

	// Crear un mock de DBConnector
	mockConnector := new(MockDBConnector)
	mockConnector.On("DBConnection", mock.Anything).Return(&gorm.DB{}, errors.New("Error connecting to db"))

	// Intentar conectar a la base de datos utilizando el mock de DBConnector
	dns := "mock_dsn" // Esto podría ser cualquier cadena ya que estamos usando un mock
	actualDB, err := db.DBConnection(mockConnector, dns)

	// Verificar que no hay errores y que la conexión se haya establecido correctamente
	assert.Error(t, err)
	assert.Nil(t, actualDB)

	// Verificar que todas las expectativas del mock se cumplieron
}
