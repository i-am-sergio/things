package probandoo

import (
	"auth-microservice/models"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Find(dest interface{}, conds ...interface{}) error {
	args := m.Called(dest, conds)
	return args.Error(0)
}

func TestGetAllUsers_Success(t *testing.T) {
	// Crea un mock de la base de datos
	mockDB := new(MockDatabase)

	// Define los resultados esperados del mock
	users := []models.User{
		{IdAuth: "123", Name: "pepe"},
		{IdAuth: "456", Name: "pepe"},
	}
	mockDB.On("Find", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*[]models.User)
		*dest = users
	})

	// Crea una instancia de DBClient con el mock de base de datos
	dbClient := &DBClient{DB: mockDB}

	// Ejecuta la función GetAllUsers
	returnedUsers, statusCode := dbClient.GetAllUsers()

	// Verifica que la función devuelva los usuarios correctos y el código de estado correcto
	assert.Equal(t, users, returnedUsers)
	assert.Equal(t, http.StatusOK, statusCode)

	// Verifica que se llamó al método Find en el mock con los argumentos esperados
	mockDB.AssertExpectations(t)
}

func TestGetAllUsers_Error(t *testing.T) {
	// Crea un mock de la base de datos
	mockDB := new(MockDatabase)

	// Define un error simulado para el método Find
	expectedError := errors.New("database error")
	mockDB.On("Find", mock.Anything, mock.Anything).Return(expectedError)

	// Crea una instancia de DBClient con el mock de base de datos
	dbClient := &DBClient{DB: mockDB}

	// Ejecuta la función GetAllUsers
	returnedUsers, statusCode := dbClient.GetAllUsers()

	// Verifica que la función devuelva nil para los usuarios y el código de estado HTTP interno del servidor
	assert.Nil(t, returnedUsers)
	assert.Equal(t, http.StatusInternalServerError, statusCode)

	// Verifica que se llamó al método Find en el mock con los argumentos esperados
	mockDB.AssertExpectations(t)
}
