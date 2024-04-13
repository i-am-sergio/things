package repository

import (
	"auth-microservice/models"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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
func TestGetAllUsers(t *testing.T) {
	// Arrange
	t.Run("Get users success", func(t *testing.T) {
		// Configura la base de datos de prueba y el repositorio de usuarios
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		implObj := NewUserRepository(db)

		// Configura las filas de datos simulados
		rows := sqlmock.NewRows([]string{"id_auth", "name", "password", "image"}).
			AddRow("1", "user1", "password1", "image1").
			AddRow("2", "user2", "password2", "image2")

		// Configura la expectativa de la consulta SQL
		expectedSQL := "SELECT (.+) FROM \"users\""
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)
		// Act
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Llama a la función GetAllUsers del repositorio de usuarios con el contexto de Echo
		users, err := implObj.GetAllUsers(c)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 2, len(users))

		// Verifica que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Get users with error", func(t *testing.T) {

		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		implObj := NewUserRepository(db)
		mock.ExpectQuery("SELECT (.+) FROM \"users\"").WillReturnError(errors.New("database error")) // Simulando un error en la base de datos
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// Act
		users, err := implObj.GetAllUsers(c)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, users)
		assert.EqualError(t, err, "database error")
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
func TestCreateUser(t *testing.T) {
	t.Run("Create User success", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		implObj := NewUserRepository(db)

		user := &models.User{
			Name:   "Juan",
			IdAuth: "1",
		}

		expectedSQL := "INSERT INTO \"users\""
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(user.Name, user.IdAuth, "", "", "", "", "USER").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Act
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		createdUser, err := implObj.CreateUser(c, user)

		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, user, createdUser)
		assert.Equal(t, user.Name, createdUser.Name)
	})

	t.Run("Create user error", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		implObj := NewUserRepository(db)

		user := &models.User{
			Name:   "Juan",
			IdAuth: "1",
		}

		expectedSQL := "INSERT INTO \"users\""
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(user.Name, user.IdAuth, "", "", "", "", "USER").
			WillReturnError(errors.New("database error")) // Simulando un error en la base de datos
		mock.ExpectRollback() // Se espera un rollback debido al error

		// Act
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		createdUser, err := implObj.CreateUser(c, user)

		// Assert
		assert.NotNil(t, err) // Se espera un error
		assert.Nil(t, createdUser)
		assert.EqualError(t, err, "database error") // Se espera que el error sea el esperado
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
func TestGetUserByIdAuth(t *testing.T) {
	t.Run("Get user by ID", func(t *testing.T) {

		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		implObj := NewUserRepository(db)

		userID := "1"
		expectedUser := &models.User{
			IdAuth:    userID,
			Name:      "John Doe",
			Password:  "john_doe",
			Image:     "password",
			Email:     "john@example.com",
			Ubication: "New York",
			Role:      "USER",
		}

		rows := sqlmock.NewRows([]string{"id_auth", "name", "password", "image", "email", "ubication", "role"}).
			AddRow(expectedUser.IdAuth, expectedUser.Name, expectedUser.Password, expectedUser.Image, expectedUser.Email, expectedUser.Ubication, expectedUser.Role)

		expectedSQL := "SELECT (.+) FROM \"users\" WHERE id_auth = ?"
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

		// Act
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users/"+userID, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		resultUser, err := implObj.GetUserByIdAuth(c, userID)

		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, resultUser)
		assert.Equal(t, expectedUser, resultUser)
	})

	t.Run("Get user by ID not found", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		implObj := NewUserRepository(db)
		users := sqlmock.NewRows([]string{"id_auth", "full_name", "user_name", "password"})

		expectedSQL := "SELECT (.+) FROM \"users\" WHERE id_auth =?"
		mock.ExpectQuery(expectedSQL).WillReturnRows(users)

		// Act
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		_, res := implObj.GetUserByIdAuth(c, "2")
		assert.True(t, errors.Is(res, gorm.ErrRecordNotFound))
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
func TestUpdateUser_Success(t *testing.T) {
	// Arrange
	t.Run("Update user", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		updatedUser := &models.User{
			IdAuth:   "1",
			Name:     "Updated Name",
			Password: "updated_password",
			Image:    "updated_image",
		}
		implObj := NewUserRepository(db)
		updUserSQL := "UPDATE \"users\" SET"
		mock.ExpectBegin()
		mock.ExpectExec(updUserSQL).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Act
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		_, err := implObj.UpdateUser(c, "1", updatedUser)

		// Assert
		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Update user error", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		implObj := NewUserRepository(db)

		// Datos del usuario actualizado
		updatedUser := &models.User{
			IdAuth:    "1",
			Name:      "Updated Name",
			Password:  "updated_password",
			Email:     "",
			Image:     "",
			Ubication: "",
			Role:      models.RoleUser,
		}

		// Expectativas de la transacción
		mock.ExpectBegin()

		// Expectativa de la ejecución de la consulta UPDATE
		mock.ExpectExec("UPDATE \"users\"  SET").
			WillReturnError(gorm.ErrRecordNotFound)

		// Rollback de la transacción
		mock.ExpectRollback()

		// Actuar
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		_, err := implObj.UpdateUser(c, "1", updatedUser)

		// Afirmar
		assert.Error(t, err)

		// Asegurar que se cumplieron todas las expectativas
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
func TestChangeUserRole_Success(t *testing.T) {
	// Arrange
	t.Run("Change role", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		implObj := NewUserRepository(db)

		// Datos del usuario
		id := "1"
		newRole := models.RoleAdmin

		// Expectativa de actualización del rol del usuario
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE \"users\" SET").
			WithArgs(newRole, id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Actuar
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/users/role/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		user, err := implObj.ChangeUserRole(c, id, newRole)

		// Afirmar
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, newRole, user.Role)

		// Asegurar que se cumplieron todas las expectativas
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Change role error", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		implObj := NewUserRepository(db)

		// Datos del usuario
		id := "1"
		newRole := models.RoleAdmin

		// Expectativa de actualización del rol del usuario con error
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE \"users\"  SET").
			WithArgs(newRole, id).
			WillReturnError(fmt.Errorf("database error"))
		mock.ExpectRollback()

		// Actuar
		// Crea un contexto de Echo vacío
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/users/role/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		user, err := implObj.ChangeUserRole(c, id, newRole)

		// Afirmar
		assert.Error(t, err)
		assert.Nil(t, user)

		// Asegurar que se cumplieron todas las expectativas
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
