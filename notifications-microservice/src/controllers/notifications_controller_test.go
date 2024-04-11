package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"notifications-microservice/src/controllers"
	"notifications-microservice/src/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

// TestGetNotificationByID_Success prueba la obtención de una notificación por su ID
func TestGetNotificationByID_Success(t *testing.T) {
	// GIVEN
	mockService := new(mockNotificationService)
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

func TestGetNotificationsByUserIDService(t *testing.T) {
	// GIVEN
	mockService := new(mockNotificationService)
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
