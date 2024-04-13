package db_test

import (
	"auth-microservice/db"
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// import (
// 	"auth-microservice/db"
// 	"errors"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"gorm.io/gorm"
// )

// // MockDB es un mock de *gorm.DB.
// type MockDB struct {
// 	mock.Mock
// }

// func (m *MockDB) DBConnection() (*gorm.DB, error) {
// 	args := m.Called()
// 	return args.Get(0).(*gorm.DB), args.Error(1)
// }

// func TestDBConnection(t *testing.T) {
// 	// Configurar el mock de *gorm.DB
// 	mockDB := new(MockDB)

// 	// Crear una instancia de DatabaseImpl usando el mock
// 	conn := db.NewConnection(mockDB)

// 	os.Setenv("DB_DNS", "host=localhost user=postgres password=admin dbname=users port=5432")
// 	gormDB := &gorm.DB{}
// 	mockDB.On("DBConnection").Return(gormDB, nil)
// 	_, err := conn.DBConnection()

// 	assert.Nil(t, err)
// 	// mockDB.AssertExpectations(t) // Asegura que todas las expectativas se cumplan
// }

// func TestDBConnectionError(t *testing.T) {
// 	// Configurar el mock de *gorm.DB
// 	mockDB := new(MockDB)

// 	// Crear una instancia de DatabaseImpl usando el mock
// 	conn := db.NewConnection(mockDB)

// 	os.Setenv("DB_DNS", "host=localhost user=postgres password=admin dbname=users port=54s32")
// 	gormDB := &gorm.DB{}
// 	mockDB.On("DBConnection").Return(gormDB, errors.New("failed to connect to database"))
// 	_, err := conn.DBConnection()

// 	assert.Error(t, err)
// 	// mockDB.AssertExpectations(t) // Asegura que todas las expectativas se cumplan
// }

// func TestNewConnectionWithNilDB(t *testing.T) {
// 	// Llama a NewConnection con db == nil
// 	conn := db.NewConnection(nil)

// 	// Asegúrate de que la instancia de conexión sea diferente de nil
// 	assert.NotNil(t, conn)

// 	// Intenta llamar a DBConnection en la instancia de conexión
// 	_, err := conn.DBConnection()

// 	// Verifica si se produce un error, ya que no se proporcionó ninguna base de datos
// 	assert.Error(t, err)
// }

func TestConnectDB(t *testing.T) {
	// Define el DNS de prueba para la conexión.
	testDNS := "host=localhost user=postgres password=admin dbname=users port=5432"

	// Llama a la función DBConnection con el DNS de prueba.
	dbInstance, err := db.DBConnection(testDNS)

	// Verifica si se produjeron errores al establecer la conexión.
	if err != nil {
		log.Println("Error connecting to the database:", err)
	}

	// Verifica si la instancia de la base de datos no es nula.
	assert.NotNil(t, dbInstance, "Database connection should not be nil")

	// Verifica si no se produjeron errores al conectar.
	assert.NoError(t, err, "Failed to connect to the database")

	// Intenta cerrar la conexión a la base de datos (si existe).
	if dbInstance != nil {
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		assert.NoError(t, err, "Failed to close the database connection")
	}
}

func TestConnectDBError(t *testing.T) {
	// Define el DNS de prueba para la conexión.
	testDNS := "host=localhostt user=postgres password=admin dbname=users port=5432"

	// Llama a la función DBConnection con el DNS de prueba.
	instanceDB, err := db.DBConnection(testDNS)

	// Verifica si se produjeron errores al establecer la conexión.
	if err != nil {
		log.Println("Error connecting to the database:", err)
	}

	// Verifica si la instancia de la base de datos no es nula.
	assert.NotNil(t, err)
	assert.Nil(t, instanceDB)

}
