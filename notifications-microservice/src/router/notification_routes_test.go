package router_test

import (
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

// Otras implementaciones de métodos de servicio mock...

func TestGetNotificationByIDRoute_Success(t *testing.T) {
	// GIVEN
	mockService := new(mockNotificationService)
	controller := controllers.NewNotificationController(mockService)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/notifications/123", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("notification_id")
	ctx.SetParamValues("123")

	// Define el comportamiento esperado del servicio mock
	expectedNotification := &models.NotificationModel{
		// Configura los campos de la notificación según sea necesario
	}
	mockService.On("GetNotificationByIDService", ctx, "123").Return(expectedNotification, nil)

	// WHEN
	err := controller.GetNotificationByID(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verificar que el servicio mock fue llamado como se esperaba
	mockService.AssertExpectations(t)
}

func TestGetNotificationsByUserIDRoute_Success(t *testing.T) {
	// GIVEN
	mockService := new(mockNotificationService)
	controller := controllers.NewNotificationController(mockService)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/notifications/user/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("user_id")
	ctx.SetParamValues("1")

	// Define el comportamiento esperado del servicio mock
	expectedNotifications := []models.NotificationModel{
		// Configura las notificaciones según sea necesario
	}
	mockService.On("GetNotificationsByUserIDService", ctx, "1").Return(expectedNotifications, nil)

	// WHEN
	err := controller.GetNotificationsByUserID(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verificar que el servicio mock fue llamado como se esperaba
	mockService.AssertExpectations(t)
}

func TestCreateNotificationRoute_Success(t *testing.T) {
	// GIVEN
	mockService := new(mockNotificationService)
	controller := controllers.NewNotificationController(mockService)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/notifications", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Define el comportamiento esperado del servicio mock
	notification := &models.NotificationModel{}
	mockService.On("CreateNotificationService", ctx, notification).Return(nil)

	// WHEN
	err := controller.CreateNotification(ctx)

	// THEN
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Verificar que el servicio mock fue llamado como se esperaba
	mockService.AssertExpectations(t)
}
