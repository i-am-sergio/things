package controllers_test

import (
	"auth-microservice/controllers"
	"auth-microservice/models"
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockUserService implementa la interfaz UserService para pruebas.
type MockUserService struct{}

func (m *MockUserService) GetAllUsersService(c echo.Context) ([]models.User, error) {
	return []models.User{
		{IdAuth: "1", Name: "User 1"},
		{IdAuth: "2", Name: "User 2"},
	}, nil
}

func (m *MockUserService) CreateUserService(c echo.Context, user *models.User) (*models.User, error) {
	return &models.User{
		IdAuth: "google-oauth2|113549236623362455025",
		Name:   user.Name,
	}, nil
}

func (m *MockUserService) UpdateUserService(c echo.Context, id string, user *models.User) (*models.User, error) {
	return &models.User{
		IdAuth:    id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Ubication: user.Ubication,
		Image:     user.Image,
	}, nil
}

func (m *MockUserService) ChangeUserRoleService(c echo.Context, id string, newRole models.Role) (*models.User, error) {
	return &models.User{
		IdAuth: id,
		Role:   newRole,
	}, nil
}

func (m *MockUserService) GetUserByIdAuthService(c echo.Context, id string) (*models.User, error) {
	return &models.User{
		IdAuth: "1",
		Name:   "User 1",
	}, nil
}
func TestGetAllUsersHandler(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := &MockUserService{}
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Llamar al controlador
	if assert.NoError(t, userController.GetAllUsersHandler(c)) {
		// Verificar la respuesta
		assert.Equal(t, http.StatusOK, rec.Code)

		// Obtener el cuerpo de la respuesta como bytes
		responseBody := rec.Body.Bytes()

		// Decodificar el cuerpo de la respuesta a una lista de usuarios
		var responseUsers []models.User
		err := json.Unmarshal(responseBody, &responseUsers)
		if err != nil {
			t.Fatalf("Error decoding response body: %s", err)
		}

		// Verificar el cuerpo de la respuesta
		assert.Equal(t, []models.User{
			{IdAuth: "1", Name: "User 1"},
			{IdAuth: "2", Name: "User 2"},
		}, responseUsers)
	}
}

var tokenJWT = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMzU0OTIzNjYyMzM2MjQ1NTAyNSIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjg5NzgyMSwiZXhwIjoxNzEyOTg0MjIxLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.LlJYUo_M0HRYBqBiECcIRiqtPFb9QsQauAmh8RTKaDEwXl4t2Yh9XeuPmqSuKT2cYvEV8YOPb6PLNrGA5JhnHTzt-nHY3Srn1EWhP96LefnNnMC9QEf-smmhVRbD-EeXm_yugQ_a3b2rIE6_BI229gw4ZRdJ7ewxraiuuwKfzfAuzi-BtbyxhFZ7QXF4UizZ84u3DDTP8yuk3nv6xUMMXEt1PKTNbegmKYvT_5Z7AQAoljcyHrjyqYWGGkcYIvKSjF6IAsp07qMy11-d74sCfyj2KzwK6_4XYcgkkebIFwAclNZiWtGxuwa0XJF4Z2HD60Ha9cuBFnXYWm72nqqfMA"

func TestCreateUserHandler(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := &MockUserService{}
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"Test User"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Añadir el token de acceso Bearer al encabezado Authorization
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenJWT)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Llamar al controlador
	if assert.NoError(t, userController.CreateUserHandler(c)) {
		// Verificar la respuesta
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Obtener el cuerpo de la respuesta como bytes
		responseBody := rec.Body.Bytes()

		// Decodificar el cuerpo de la respuesta a un usuario
		var createdUser models.User
		err := json.Unmarshal(responseBody, &createdUser)
		if err != nil {
			t.Fatalf("Error decoding response body: %s", err)
		}

		// Verificar el cuerpo de la respuesta
		assert.Equal(t, &models.User{IdAuth: "google-oauth2|113549236623362455025", Name: "Test User"}, &createdUser)
	}
}
func TestUpdateUserHandler(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := &MockUserService{}
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()

	// Crear un búfer de bytes para el cuerpo de la solicitud
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Agregar campos al formulario
	writer.WriteField("name", "Updated User")

	// Finalizar la escritura del formulario
	err := writer.Close()
	if err != nil {
		t.Fatalf("Error closing multipart writer: %s", err)
	}

	// Crear una solicitud HTTP multipart/form-data
	req := httptest.NewRequest(http.MethodPut, "/users/1", body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Llamar al controlador
	if assert.NoError(t, userController.UpdateUserHandler(c)) {
		// Verificar la respuesta
		assert.Equal(t, http.StatusOK, rec.Code)

		// Obtener el cuerpo de la respuesta como bytes
		responseBody := rec.Body.Bytes()

		// Decodificar el cuerpo de la respuesta a un usuario actualizado
		var updatedUser models.User
		err := json.Unmarshal(responseBody, &updatedUser)
		if err != nil {
			t.Fatalf("Error decoding response body: %s", err)
		}

		// Verificar el cuerpo de la respuesta
		assert.Equal(t, &models.User{IdAuth: "1", Name: "Updated User"}, &updatedUser)
	}
}
func TestChangeRoleHandler(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := &MockUserService{}
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/role/1", strings.NewReader(`"ADMIN"`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Llamar al controlador
	if assert.NoError(t, userController.ChangeRoleHandler(c)) {
		// Verificar la respuesta
		assert.Equal(t, http.StatusOK, rec.Code)

		// Obtener el cuerpo de la respuesta como bytes
		responseBody := rec.Body.Bytes()

		// Decodificar el cuerpo de la respuesta a un usuario actualizado con el nuevo rol
		var updatedUser models.User
		err := json.Unmarshal(responseBody, &updatedUser)
		if err != nil {
			t.Fatalf("Error decoding response body: %s", err)
		}

		// Verificar el cuerpo de la respuesta
		assert.Equal(t, models.RoleAdmin, updatedUser.Role)
	}
}

func TestGetUserHandler(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := &MockUserService{}
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Llamar al controlador
	if assert.NoError(t, userController.GetUserHandler(c)) {
		// Verificar la respuesta
		assert.Equal(t, http.StatusOK, rec.Code)

		// Obtener el cuerpo de la respuesta como bytes
		responseBody := rec.Body.Bytes()

		// Decodificar el cuerpo de la respuesta a un usuario
		var foundUser models.User
		err := json.Unmarshal(responseBody, &foundUser)
		if err != nil {
			t.Fatalf("Error decoding response body: %s", err)
		}

		// Verificar el cuerpo de la respuesta
		assert.Equal(t, &models.User{IdAuth: "1", Name: "User 1"}, &foundUser)
	}
}
