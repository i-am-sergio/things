package services

import (
	"auth-microservice/models"
	"database/sql"
	"log"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var u = &models.User{
	IdAuth:    "123",
	Name:      "Momo",
	Email:     "momo@mail.com",
	Image:     "08123456789",
	Password:  "password123",
	Ubication: "Some Location",
	Role:      "USER",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id_auth, name, email, image FROM users WHERE id_auth = \\?"

	rows := sqlmock.NewRows([]string{"id_auth", "name", "email", "image"}).
		AddRow(u.IdAuth, u.Name, u.Email, u.Image)

	mock.ExpectQuery(query).WithArgs(u.IdAuth).WillReturnRows(rows)

	user, err := repo.GetUserByIdAuth(u.IdAuth)
	assert.NotNil(t, user)
	assert.Equal(t, http.StatusOK, err)
}

func TestFindByIDError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id_auth, name, email, image FROM users WHERE id_auth = \\?"

	rows := sqlmock.NewRows([]string{"id_auth", "name", "email", "image"})

	mock.ExpectQuery(query).WithArgs(u.IdAuth).WillReturnRows(rows)

	user, err := repo.GetUserByIdAuth(u.IdAuth)
	assert.Empty(t, user)
	assert.Equal(t, http.StatusNotFound, err)
}

func TestFind(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id_auth, name, email, image FROM users"

	rows := sqlmock.NewRows([]string{"id_auth", "name", "email", "image"}).
		AddRow(u.IdAuth, u.Name, u.Email, u.Image)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.GetAllUsers()
	assert.NotEmpty(t, users)
	assert.Equal(t, http.StatusOK, err)
	assert.Len(t, users, 1)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO users \\(id_auth, name, email, password, image, ubication, role\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.IdAuth, u.Name, u.Email, u.Password, u.Image, u.Ubication, u.Role).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := repo.CreateUser(u)
	assert.Equal(t, http.StatusCreated, err)

}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE users SET name = \\?, email = \\?, password = \\?, image = \\?, ubication = \\?, role = \\? WHERE id_auth = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.Password, u.Image, u.Ubication, u.Role, u.IdAuth).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := repo.UpdateUser(u.IdAuth, u)
	assert.Equal(t, http.StatusOK, err)
}
