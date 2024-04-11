package services

import (
	"auth-microservice/models"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) First(dest interface{}, conds ...interface{}) error {
	args := m.Called(append([]interface{}{dest}, conds...)...)
	return args.Error(0)
}

func (m *MockDatabase) Save(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDatabase) Create(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDatabase) Where(query interface{}, args ...interface{}) Database {
	allArgs := append([]interface{}{query}, args...)
	result := m.Called(allArgs...)
	return result.Get(0).(Database)
}

func (m *MockDatabase) Update(attrs ...interface{}) error {
	args := m.Called(attrs...)
	return args.Error(0)
}
func (m *MockDatabase) Find(dest interface{}, conds ...interface{}) error {
	args := m.Called(append([]interface{}{dest}, conds...)...)
	return args.Error(0)
}

func (m *MockDatabase) FindPreloaded(relation string, dest interface{}, conds ...interface{}) error {
	args := m.Called(append([]interface{}{relation, dest}, conds...)...)
	return args.Error(0)
}

func (m *MockDatabase) Delete(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDatabase) DeleteByID(model interface{}, id interface{}) error {
	args := m.Called(model, id)
	return args.Error(0)
}

func (m *MockDatabase) Model(value interface{}) Database {
	args := m.Called(value)
	return args.Get(0).(Database)
}

// }
// type UserService struct {
// 	repository RepositoryFunc
// }

// func NewUserService(repo RepositoryFunc) *UserService {
// 	return &UserService{repository: repo}
// }

// func (s *UserService) GetAllUsers() ([]models.User, int) {
// 	return s.repository.GetAllUsers()
// }

// func NewMock() *MockRepository {
// 	return &MockRepository{}
// }

// func (m *MockRepository) GetAllUsers() ([]models.User, int) {

// 	args := m.Called()
// 	if args.Get(0) == nil {
// 		return nil, args.Int(1)
// 	}
// 	return args.Get(0).([]models.User), args.Int(1)
// }

// func (m *MockRepository) GetUserByIdAuth(idAuth string) (*models.User, int) {
// 	args := m.Called(idAuth)
// 	if args.Get(0) == nil {
// 		return nil, args.Int(1)
// 	}
// 	return args.Get(0).(*models.User), args.Int(1)
// }

// func (m *MockRepository) CreateUser(user *models.User) (*models.User, int) {
// 	args := m.Called(user)
// 	if args.Get(0) == nil {
// 		return nil, args.Int(1)
// 	}
// 	return args.Get(0).(*models.User), args.Int(1)
// }

// func (m *MockRepository) UpdateUser(id string, user *models.User) (*models.User, int) {
// 	args := m.Called(id, user)
// 	if args.Get(0) == nil {
// 		return nil, args.Int(1)
// 	}
// 	return args.Get(0).(*models.User), args.Int(1)
// }

//	func (m *MockRepository) ChangeUserRole(id string, newRole models.Role) (*models.User, int) {
//		args := m.Called(id, newRole)
//		if args.Get(0) == nil {
//			return nil, args.Int(1)
//		}
//		return args.Get(0).(*models.User), args.Int(1)
//	}
func TestGetAllUsers(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
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
	})

	t.Run("Error", func(t *testing.T) {
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
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Crea un mock de la base de datos
		mockDB := new(MockDatabase)

		// Crea una instancia de usuario para crear
		userToCreate := &models.User{IdAuth: "123", Name: "pepe"}

		// Define los resultados esperados del mock para Create
		mockDB.On("Create", userToCreate).Return(nil)

		// Crea una instancia de DBClient con el mock de base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecuta la función CreateUser
		createdUser, statusCode := dbClient.CreateUser(userToCreate)

		// Verifica que la función devuelva el usuario creado y el código de estado HTTP OK
		assert.Equal(t, userToCreate, createdUser)
		assert.Equal(t, http.StatusOK, statusCode)

		// Verifica que se llamó al método Create en el mock con los argumentos esperados
		mockDB.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// Crea un mock de la base de datos
		mockDB := new(MockDatabase)

		// Crea una instancia de usuario para crear
		userToCreate := &models.User{IdAuth: "123", Name: "pepe"}

		// Define un error simulado para el método Create
		expectedError := errors.New("database error")
		mockDB.On("Create", userToCreate).Return(expectedError)

		// Crea una instancia de DBClient con el mock de base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecuta la función CreateUser
		createdUser, statusCode := dbClient.CreateUser(userToCreate)

		// Verifica que la función devuelva nil para el usuario y el código de estado HTTP interno del servidor
		assert.Nil(t, createdUser)
		assert.Equal(t, http.StatusInternalServerError, statusCode)

		// Verifica que se llamó al método Create en el mock con los argumentos esperados
		mockDB.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Usuario actualizado
		updatedUser := &models.User{IdAuth: "123", Name: "Updated"}

		// Mock para la búsqueda del usuario
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("First", mock.AnythingOfType("*models.User")).Return(nil)

		// Mock para guardar los cambios
		mockDB.On("Save", updatedUser).Return(nil)

		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función UpdateUser
		updatedUserResult, statusCode := dbClient.UpdateUser("123", updatedUser)

		// Verificar que el usuario actualizado y el código de estado son correctos
		assert.Equal(t, updatedUser, updatedUserResult)
		assert.Equal(t, http.StatusOK, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})

	t.Run("Error_Find", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Mock para la búsqueda del usuario que devuelve un error
		expectedError := errors.New("database error")
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("First", mock.AnythingOfType("*models.User")).Return(expectedError)

		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función UpdateUser
		updatedUserResult, statusCode := dbClient.UpdateUser("123", &models.User{})

		// Verificar que se devuelva un código de estado HTTP interno del servidor
		assert.Nil(t, updatedUserResult)
		assert.Equal(t, http.StatusInternalServerError, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})

	t.Run("Error_Save", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Usuario actualizado
		updatedUser := &models.User{IdAuth: "123", Name: "Updated"}

		// Mock para la búsqueda del usuario
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("First", mock.AnythingOfType("*models.User")).Return(nil)

		// Mock para guardar los cambios que devuelve un error
		expectedError := errors.New("database error")
		mockDB.On("Save", updatedUser).Return(expectedError)

		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función UpdateUser
		updatedUserResult, statusCode := dbClient.UpdateUser("123", updatedUser)

		// Verificar que se devuelva un código de estado HTTP interno del servidor
		assert.Nil(t, updatedUserResult)
		assert.Equal(t, http.StatusInternalServerError, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})
}
func TestGetUserByIdAuth(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Usuario encontrado en la base de datos
		user := &models.User{IdAuth: "123", Name: "John"}

		// Mock para la consulta que devuelve el usuario encontrado
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("First", mock.AnythingOfType("*models.User")).Return(nil).Run(func(args mock.Arguments) {
			dest := args.Get(0).(*models.User)
			*dest = *user
		})
		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función GetUserByIdAuth
		userResult, statusCode := dbClient.GetUserByIdAuth("123")

		// Verificar que el usuario encontrado y el código de estado son correctos
		assert.Equal(t, user, userResult)
		assert.Equal(t, http.StatusOK, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Mock para la consulta que no encuentra el usuario
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("First", mock.AnythingOfType("*models.User"), mock.Anything).Return(gorm.ErrRecordNotFound)

		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función GetUserByIdAuth
		userResult, statusCode := dbClient.GetUserByIdAuth("123")

		// Verificar que se devuelva el código de estado HTTP de recurso no encontrado
		assert.Nil(t, userResult)
		assert.Equal(t, http.StatusNotFound, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Mock para la consulta que devuelve un error
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("First", mock.AnythingOfType("*models.User"), mock.Anything).Return(errors.New("database error"))

		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función GetUserByIdAuth
		userResult, statusCode := dbClient.GetUserByIdAuth("123")

		// Verificar que se devuelva el código de estado HTTP interno del servidor
		assert.Nil(t, userResult)
		assert.Equal(t, http.StatusInternalServerError, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})
}

func TestChangeUserRole(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Nuevo rol para el usuario
		newRole := models.RoleAdmin

		// Mock para la consulta que actualiza el rol del usuario
		mockDB.On("Model", &models.User{}).Return(mockDB)
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("Update", "role", newRole).Return(nil)

		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función ChangeUserRole
		userResult, statusCode := dbClient.ChangeUserRole("123", newRole)

		// Verificar que el usuario actualizado y el código de estado son correctos
		expectedUser := &models.User{IdAuth: "123", Role: newRole}
		assert.Equal(t, expectedUser, userResult)
		assert.Equal(t, http.StatusOK, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// Mock de la base de datos
		mockDB := new(MockDatabase)

		// Nuevo rol para el usuario
		newRole := models.RoleAdmin

		// Error simulado al actualizar el rol del usuario
		expectedError := errors.New("database error")
		mockDB.On("Model", &models.User{}).Return(mockDB)
		mockDB.On("Where", "id_auth = ?", "123").Return(mockDB)
		mockDB.On("Update", "role", newRole).Return(expectedError)

		// Crear instancia de DBClient con el mock de la base de datos
		dbClient := &DBClient{DB: mockDB}

		// Ejecutar la función ChangeUserRole
		userResult, statusCode := dbClient.ChangeUserRole("123", newRole)

		// Verificar que se devuelva nil para el usuario y el código de estado HTTP interno del servidor
		assert.Nil(t, userResult)
		assert.Equal(t, http.StatusInternalServerError, statusCode)

		// Verificar que los métodos del mock se llamaron con los argumentos correctos
		mockDB.AssertExpectations(t)
	})
}

// func TestGetUserById(t *testing.T) {

// 	var userForTest = &models.User{
// 		IdAuth:    "123",
// 		Name:      "Momo",
// 		Email:     "momo@mail.com",
// 		Image:     "08123456789",
// 		Password:  "password123",
// 		Ubication: "Some Location",
// 		Role:      "USER",
// 	}
// 	t.Run("Find by id that exist", func(t *testing.T) {
// 		repo := NewMock()

// 		repo.On("GetUserByIdAuth", userForTest.IdAuth).Return(userForTest, http.StatusOK)

// 		user, errCode := repo.GetUserByIdAuth(userForTest.IdAuth)

// 		assert.NotNil(t, user)
// 		assert.Equal(t, http.StatusOK, errCode)
// 	})
// 	t.Run("Find by id not found", func(t *testing.T) {
// 		repo := NewMock()
// 		repo.On("GetUserByIdAuth", userForTest.IdAuth).Return(nil, http.StatusNotFound)

// 		user, errCode := repo.GetUserByIdAuth(userForTest.IdAuth)
// 		assert.Empty(t, user)
// 		assert.Equal(t, http.StatusNotFound, errCode)
// 	})

// 	t.Run("Find by id with internal server error", func(t *testing.T) {
// 		repo := NewMock()

// 		repo.On("GetUserByIdAuth", userForTest.IdAuth).Return(nil, http.StatusInternalServerError)

// 		user, errCode := repo.GetUserByIdAuth(userForTest.IdAuth)

// 		assert.Empty(t, user)
// 		assert.Equal(t, http.StatusInternalServerError, errCode)
// 	})

// }

// func TestCreateUser(t *testing.T) {
// 	t.Run("Success", func(t *testing.T) {
// 		// Crear un nuevo usuario
// 		user := &models.User{
// 			IdAuth:    "1",
// 			Name:      "User 1",
// 			Email:     "user1@example.com",
// 			Image:     "image1",
// 			Password:  "password",
// 			Ubication: "Location",
// 			Role:      "USER",
// 		}
// 		// Crear un nuevo mock repository
// 		repo := NewMock()

// 		// Configurar el comportamiento esperado del mock para CreateUser en caso de éxito
// 		repo.On("CreateUser", user).Return(user, http.StatusCreated)

// 		// Llamar a la función que estamos probando
// 		createdUser, status := repo.CreateUser(user)

// 		// Verificar los resultados
// 		assert.NotNil(t, createdUser)
// 		assert.Equal(t, http.StatusCreated, status)
// 	})

// 	t.Run("InternalServerError", func(t *testing.T) {
// 		// Crear un nuevo usuario
// 		user := &models.User{
// 			IdAuth:    "1",
// 			Name:      "User 1",
// 			Email:     "user1@example.com",
// 			Image:     "image1",
// 			Password:  "password",
// 			Ubication: "Location",
// 			Role:      "USER",
// 		}

// 		// Crear un nuevo mock repository
// 		repo := NewMock()

// 		// Configurar el comportamiento esperado del mock para CreateUser en caso de error interno del servidor
// 		repo.On("CreateUser", user).Return(nil, http.StatusInternalServerError)

// 		// Llamar a la función que estamos probando
// 		createdUser, status := repo.CreateUser(user)

// 		// Verificar que se devuelva el código de estado HTTP 500 (InternalServerError)
// 		assert.Nil(t, createdUser)
// 		assert.Equal(t, http.StatusInternalServerError, status)
// 	})
// }

// func TestUpdateUser(t *testing.T) {
// 	t.Run("Success", func(t *testing.T) {
// 		// Crear un usuario actualizado
// 		updatedUser := &models.User{
// 			IdAuth:    "1",
// 			Name:      "Updated User",
// 			Email:     "updateduser@example.com",
// 			Image:     "updatedimage",
// 			Password:  "updatedpassword",
// 			Ubication: "Updated Location",
// 			Role:      "UPDATED_ROLE",
// 		}

// 		// Crear un nuevo mock repository
// 		repo := NewMock()

// 		// Configurar el comportamiento esperado del mock para UpdateUser en caso de éxito
// 		repo.On("UpdateUser", "1", updatedUser).Return(updatedUser, http.StatusOK)

// 		// Llamar a la función que estamos probando
// 		updated, status := repo.UpdateUser("1", updatedUser)

// 		// Verificar los resultados
// 		assert.NotNil(t, updated)
// 		assert.Equal(t, http.StatusOK, status)
// 	})

// 	t.Run("InternalServerError", func(t *testing.T) {
// 		// Crear un usuario actualizado
// 		updatedUser := &models.User{
// 			IdAuth:    "1",
// 			Name:      "Updated User",
// 			Email:     "updateduser@example.com",
// 			Image:     "updatedimage",
// 			Password:  "updatedpassword",
// 			Ubication: "Updated Location",
// 			Role:      "UPDATED_ROLE",
// 		}
// 		// Crear un nuevo mock repository
// 		repo := NewMock()

// 		// Configurar el comportamiento esperado del mock para UpdateUser en caso de error interno del servidor
// 		repo.On("UpdateUser", "1", updatedUser).Return(nil, http.StatusInternalServerError)

// 		// Llamar a la función que estamos probando
// 		updated, status := repo.UpdateUser("1", updatedUser)

// 		// Verificar que se devuelva el código de estado HTTP 500 (InternalServerError)
// 		assert.Nil(t, updated)
// 		assert.Equal(t, http.StatusInternalServerError, status)
// 	})
// }

// func TestChangeUserRole(t *testing.T) {
// 	t.Run("Success", func(t *testing.T) {
// 		// Crear un nuevo mock repository
// 		repo := NewMock()

// 		// Definir el nuevo rol
// 		newRole := models.RoleUser
// 		// Configurar el comportamiento esperado del mock para ChangeUserRole en caso de éxito
// 		repo.On("ChangeUserRole", "1", newRole).Return(&models.User{IdAuth: "1", Role: newRole}, http.StatusOK)

// 		// Llamar a la función que estamos probando
// 		user, status := repo.ChangeUserRole("1", newRole)

// 		// Verificar los resultados
// 		assert.NotNil(t, user)
// 		assert.Equal(t, http.StatusOK, status)
// 	})

// 	t.Run("InternalServerError", func(t *testing.T) {
// 		// Crear un nuevo mock repository
// 		repo := NewMock()

// 		// Definir el nuevo rol
// 		newRole := models.RoleUser

// 		// Configurar el comportamiento esperado del mock para ChangeUserRole en caso de error interno del servidor
// 		repo.On("ChangeUserRole", "1", newRole).Return(nil, http.StatusInternalServerError)

// 		// Llamar a la función que estamos probando
// 		user, status := repo.ChangeUserRole("1", newRole)

// 		// Verificar que se devuelva el código de estado HTTP 500 (InternalServerError)
// 		assert.Nil(t, user)
// 		assert.Equal(t, http.StatusInternalServerError, status)
// 	})
// }
