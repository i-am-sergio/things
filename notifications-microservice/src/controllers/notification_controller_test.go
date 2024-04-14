package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"notifications-microservice/src/controllers"
	"notifications-microservice/src/mocks"
	"notifications-microservice/src/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetNotificationByID_Success(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	notificationID := "123"
	expectedNotification := &models.NotificationModel{
		Id:      "123",
		UserID:  "1",
		Title:   "Test Title",
		Message: "Test Message",
		IsRead:  false,
	}

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("notification_id")
	ctx.SetParamValues(notificationID)

	// Define the expected behavior of the mock
	mockService.On("GetNotificationByIDService", ctx, notificationID).Return(expectedNotification, nil)

	// WHEN
	err := controller.GetNotificationByID(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Get the response body as a string and deserialize it
	responseBody := rec.Body.String()
	var responseNotification models.NotificationModel
	if err := json.Unmarshal([]byte(responseBody), &responseNotification); err != nil {
		t.Fatalf("Error deserializando la respuesta JSON: %v", err)
	}
	// Verify that the response matches the expected notification
	assert.Equal(t, expectedNotification, &responseNotification)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

func TestGetNotificationByID_NotFound(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	notificationID := "123"
	expectedError := errors.New("Notification not found")

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("notification_id")
	ctx.SetParamValues(notificationID)

	// Define the expected behavior of the mock
	mockService.On("GetNotificationByIDService", ctx, notificationID).Return(&models.NotificationModel{}, expectedError)

	// WHEN
	err := controller.GetNotificationByID(ctx)

	// THEN
	assert.NoError(t, err)                         // No debería haber error en el controlador
	assert.Equal(t, http.StatusNotFound, rec.Code) // Debería devolver un código 404

	expectedResponseBody := `{"error":"Notification not found"}`
	actualResponseBody := rec.Body.String()

	// Comparar los cuerpos de respuesta decodificados de JSON como cadenas
	assert.JSONEq(t, expectedResponseBody, actualResponseBody)

	// Verificar que el servicio mock fue llamado como se esperaba
	mockService.AssertExpectations(t)
}

func TestGetNotificationsByUserIDService(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	userID := "1"
	expectedNotifications := []models.NotificationModel{
		{
			Id:      "1",
			UserID:  "1",
			Title:   "Test Title",
			Message: "Test Message",
			IsRead:  false,
		},
		{
			Id:      "2",
			UserID:  "1",
			Title:   "Test Title 2",
			Message: "Test Message 2",
			IsRead:  false,
		},
	}

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("user_id")
	ctx.SetParamValues(userID)

	// Define the expected behavior of the mock
	mockService.On("GetNotificationsByUserIDService", ctx, userID).Return(expectedNotifications, nil)

	// WHEN
	err := controller.GetNotificationsByUserID(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Get the response body as a string and deserialize it
	responseBody := rec.Body.String()
	var responseNotifications []models.NotificationModel
	if err := json.Unmarshal([]byte(responseBody), &responseNotifications); err != nil {
		t.Fatalf("Error deserializando la respuesta JSON: %v", err)
	}
	// Verify that the response matches the expected notifications
	assert.Equal(t, expectedNotifications, responseNotifications)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

func TestGetNotificationsByUserID_NotFound(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	userID := "1"
	expectedError := errors.New("Failed to get notifications")

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("user_id")
	ctx.SetParamValues(userID)

	// Define the expected behavior of the mock
	mockService.On("GetNotificationsByUserIDService", ctx, userID).Return([]models.NotificationModel{}, expectedError)

	// WHEN
	err := controller.GetNotificationsByUserID(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	expectedResponseBody := `{"error":"Failed to get notifications"}`
	actualResponseBody := rec.Body.String()

	// Comparar los cuerpos de respuesta decodificados de JSON como cadenas
	assert.JSONEq(t, expectedResponseBody, actualResponseBody)

	// Verificar que el servicio mock fue llamado como se esperaba
	mockService.AssertExpectations(t)

}

func TestCreateNotification_Success(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	notification := &models.NotificationModel{
		UserID:  "1",
		Title:   "Test Title",
		Message: "Test Message",
		IsRead:  false,
	}
	notificationJSON, _ := json.Marshal(notification)

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(notificationJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	// ctx.SetPath("/")

	// Define the expected behavior of the mock
	mockService.On("CreateNotificationService", ctx, notification).Return(nil)

	// WHEN
	err := controller.CreateNotification(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

func TestCreateNotification_BadRequest(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	notification := &models.NotificationModel{
		UserID:  "1",
		Title:   "Test Title",
		Message: "Test Message",
		IsRead:  false,
	}
	notificationJSON, _ := json.Marshal(notification)

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(notificationJSON))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// WHEN
	err := controller.CreateNotification(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	// Verify that the mock was not called
	mockService.AssertNotCalled(t, "CreateNotificationService", ctx, notification)
}

func TestCreateNotification_InternalServerError(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	notification := &models.NotificationModel{
		UserID:  "1",
		Title:   "Test Title",
		Message: "Test Message",
		IsRead:  false,
	}
	notificationJSON, _ := json.Marshal(notification)

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(notificationJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Define the expected behavior of the mock
	mockService.On("CreateNotificationService", ctx, notification).Return(errors.New("Failed to create notification"))

	// WHEN
	err := controller.CreateNotification(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

func TestMarkAsRead_Success(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	notificationID := "123"

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("notification_id")
	ctx.SetParamValues(notificationID)

	// Define the expected behavior of the mock
	mockService.On("MarkAsReadService", ctx, notificationID).Return(nil)

	// WHEN
	err := controller.MarkAsRead(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

func TestMarkAsRead_InternalServerError(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	notificationID := "123"

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("notification_id")
	ctx.SetParamValues(notificationID)

	// Define the expected behavior of the mock
	mockService.On("MarkAsReadService", ctx, notificationID).Return(errors.New("Failed to mark notification as read"))

	// WHEN
	err := controller.MarkAsRead(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

func TestMarkAllAsRead_Success(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	userID := "1"

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("user_id")
	ctx.SetParamValues(userID)

	// Define the expected behavior of the mock
	mockService.On("MarkAllAsReadService", ctx, userID).Return(nil)

	// WHEN
	err := controller.MarkAllAsRead(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

func TestMarkAllAsRead_InternalServerError(t *testing.T) {
	// GIVEN
	mockService := new(mocks.NotificationService)
	userID := "1"

	controller := controllers.NewNotificationController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("user_id")
	ctx.SetParamValues(userID)

	// Define the expected behavior of the mock
	mockService.On("MarkAllAsReadService", ctx, userID).Return(errors.New("Failed to mark all notifications as read"))

	// WHEN
	err := controller.MarkAllAsRead(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	// Verify that the mock was called as expected
	mockService.AssertExpectations(t)
}

/*
// Use this mocks in the test file if not using mockery generator

type mockNotificationService struct {
	mock.Mock
}

func (m *mockNotificationService) GetNotificationByIDService(c echo.Context, id string) (*models.NotificationModel, error) {
	args := m.Called(c, id)
	return args.Get(0).(*models.NotificationModel), args.Error(1)
}

func (m *mockNotificationService) GetNotificationsByUserIDService(c echo.Context, id string) ([]models.NotificationModel, error) {
	args := m.Called(c, id)
	return args.Get(0).([]models.NotificationModel), args.Error(1)
}

func (m *mockNotificationService) CreateNotificationService(c echo.Context, notification *models.NotificationModel) error {
	args := m.Called(c, notification)
	return args.Error(0)
}

func (m *mockNotificationService) MarkAsReadService(c echo.Context, id string) error {
	args := m.Called(c, id)
	return args.Error(0)
}

func (m *mockNotificationService) MarkAllAsReadService(c echo.Context, id string) error {
	args := m.Called(c, id)
	return args.Error(0)
}
*/
