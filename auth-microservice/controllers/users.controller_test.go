package controllers_test

import (
	"auth-microservice/controllers"
	"auth-microservice/models"
	"auth-microservice/utils"
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService implementa la interfaz UserService para pruebas.
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUsersService(c echo.Context) ([]models.User, error) {
	args := m.Called(c)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) CreateUserService(c echo.Context, user *models.User) (*models.User, error) {
	args := m.Called(c, user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) UpdateUserService(c echo.Context, id string, user *models.User) (*models.User, error) {
	args := m.Called(c, id, user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) ChangeUserRoleService(c echo.Context, id string, newRole models.Role) (*models.User, error) {
	args := m.Called(c, id, newRole)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserByIdAuthService(c echo.Context, id string) (*models.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(*models.User), args.Error(1)
}
func TestGetAllUsersHandler(t *testing.T) {

	mocks := new(MockUserService)
	usercon := controllers.NewUserController(mocks)
	expectUsers := []models.User{
		{IdAuth: "1", Name: "User 1"},
		{IdAuth: "2", Name: "User 2"}}
	mocks.On("GetAllUsersService", mock.Anything).Return(expectUsers, nil)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, usercon.GetAllUsersHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		responseBody := rec.Body.Bytes()

		var responseUsers []models.User
		err := json.Unmarshal(responseBody, &responseUsers)
		if err != nil {
			t.Fatalf("Error decoding response body: %s", err)
		}

		assert.Equal(t, expectUsers, responseUsers)
	}
	mocks.AssertExpectations(t)
}
func TestGetAllUsersHandler_Error(t *testing.T) {
	mockService := new(MockUserService)

	userController := controllers.NewUserController(mockService)

	mockService.On("GetAllUsersService", mock.Anything).Return([]models.User{}, errors.New("Internal server error"))

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController.GetAllUsersHandler(c)

	mockService.AssertCalled(t, "GetAllUsersService", mock.Anything)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertExpectations(t)

}

var tokenJWT = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IktVWUlsME5QWlp1bXVuR2JEcVFpTSJ9.eyJpc3MiOiJodHRwczovL2Rldi1zNThybXZ0anczbnc4eWs0LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMzU0OTIzNjYyMzM2MjQ1NTAyNSIsImF1ZCI6WyJodHRwczovL2FwaS1nb2xhbmctdGVzdCIsImh0dHBzOi8vZGV2LXM1OHJtdnRqdzNudzh5azQudXMuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTcxMjg5NzgyMSwiZXhwIjoxNzEyOTg0MjIxLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiYXpwIjoiUFFTNjNuSjY5N1VpTU01dlBFYlF5Q28yTjJ1WFZBVTMifQ.LlJYUo_M0HRYBqBiECcIRiqtPFb9QsQauAmh8RTKaDEwXl4t2Yh9XeuPmqSuKT2cYvEV8YOPb6PLNrGA5JhnHTzt-nHY3Srn1EWhP96LefnNnMC9QEf-smmhVRbD-EeXm_yugQ_a3b2rIE6_BI229gw4ZRdJ7ewxraiuuwKfzfAuzi-BtbyxhFZ7QXF4UizZ84u3DDTP8yuk3nv6xUMMXEt1PKTNbegmKYvT_5Z7AQAoljcyHrjyqYWGGkcYIvKSjF6IAsp07qMy11-d74sCfyj2KzwK6_4XYcgkkebIFwAclNZiWtGxuwa0XJF4Z2HD60Ha9cuBFnXYWm72nqqfMA"

func TestCreateUserHandler(t *testing.T) {
	mocks := new(MockUserService)
	usercon := controllers.NewUserController(mocks)
	expectUser := &models.User{IdAuth: "google-oauth2|113549236623362455025", Name: "Test User"}

	mocks.On("CreateUserService", mock.Anything, mock.Anything).Return(expectUser, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"Test User"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenJWT)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, usercon.CreateUserHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
	mocks.AssertExpectations(t)
}

func TestCreateUserHandlerError(t *testing.T) {
	mocks := new(MockUserService)
	usercon := controllers.NewUserController(mocks)

	mocks.On("CreateUserService", mock.Anything, mock.Anything).Return(&models.User{}, errors.New("Internal server error"))

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"Test User"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenJWT)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	usercon.CreateUserHandler(c)

	mocks.AssertCalled(t, "CreateUserService", mock.Anything, mock.Anything)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	mocks.AssertExpectations(t)
}

func TestUpdateUserHandler(t *testing.T) {
	mocks := new(MockUserService)
	usercon := controllers.NewUserController(mocks)
	expectUser := &models.User{}
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	utils.Init("../.env")
	writer.WriteField("name", "Updated Name")
	writer.WriteField("password", "new_password")
	writer.WriteField("email", "new_email@example.com")
	writer.WriteField("ubication", "New Ubication")
	writer.CreateFormFile("image", "https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Sunflower_from_Silesia2.jpg/800px-Sunflower_from_Silesia2.jpg")
	writer.Close()
	mocks.On("UpdateUserService", mock.Anything, mock.Anything, mock.Anything).Return(expectUser, nil)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/1", &requestBody)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, usercon.UpdateUserHandler(c)) {
		// Verificar el código de estado de la respuesta
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	mocks.AssertCalled(t, "UpdateUserService", mock.Anything, "1", mock.Anything)
	mocks.AssertExpectations(t)
}
func TestUpdateUserHandlerError(t *testing.T) {
	mocks := new(MockUserService)
	usercon := controllers.NewUserController(mocks)

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	writer.Close()
	mocks.On("UpdateUserService", mock.Anything, mock.Anything, mock.Anything).Return(&models.User{}, errors.New("Mensaje oculto v:"))
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/1", &requestBody)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	usercon.UpdateUserHandler(c)
	mocks.AssertCalled(t, "UpdateUserService", mock.Anything, "1", mock.Anything)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mocks.AssertExpectations(t)
}
func TestChangeRoleHandler(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := new(MockUserService)
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/role/1", strings.NewReader(`"ADMIN"`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Configurar el comportamiento esperado del mock
	expectedUser := &models.User{Role: models.RoleAdmin}
	mockService.On("ChangeUserRoleService", mock.Anything, "1", models.RoleAdmin).Return(expectedUser, nil)

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
		assert.Equal(t, expectedUser, &updatedUser)
	}
}

func TestChangeRoleHandler_Error(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := new(MockUserService)
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/role/1", strings.NewReader(`"ADMIN"`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Configurar el comportamiento esperado del mock para devolver un error
	mockService.On("ChangeUserRoleService", mock.Anything, "1", models.RoleAdmin).Return(&models.User{}, errors.New("User not found"))

	// Llamar al controlador
	userController.ChangeRoleHandler(c)

	// Verificar que se produce un error y se devuelve un código de estado http.StatusNotFound
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockService.AssertExpectations(t)
}
func TestGetUserHandler(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := new(MockUserService)
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Configurar el usuario esperado y el error para el mock
	expectUser := &models.User{IdAuth: "1", Name: "Test User"}
	mockService.On("GetUserByIdAuthService", mock.Anything, "1").Return(expectUser, nil)

	// Llamar al controlador
	err := userController.GetUserHandler(c)

	// Verificar que no hay errores y se devuelve un código de estado http.StatusOK
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verificar el cuerpo de la respuesta
	mockService.AssertExpectations(t)
}

func TestGetUserHandler_Error(t *testing.T) {
	// Configurar el controlador con el servicio mock
	mockService := new(MockUserService)
	userController := controllers.NewUserController(mockService)

	// Configurar el enrutador Echo para la prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Configurar el error para el mock
	mockService.On("GetUserByIdAuthService", mock.Anything, "1").Return(&models.User{}, errors.New("User not found"))

	// Llamar al controlador
	userController.GetUserHandler(c)

	// Verificar que se devuelve un error y se devuelve un código de estado http.StatusNotFound
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockService.AssertExpectations(t)
}
