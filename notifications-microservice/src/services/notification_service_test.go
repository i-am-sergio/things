package services_test

import (
	"notifications-microservice/src/models"
	"notifications-microservice/src/services"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) GetNotificationByID(ctx echo.Context, id string) (*models.NotificationModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.NotificationModel), args.Error(1)
}

func (m *MockNotificationRepository) GetNotificationsByUserID(ctx echo.Context, id string) ([]models.NotificationModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]models.NotificationModel), args.Error(1)
}

func (m *MockNotificationRepository) CreateNotification(ctx echo.Context, notification *models.NotificationModel) error {
	args := m.Called(ctx, notification)
	return args.Error(0)
}

func (m *MockNotificationRepository) MarkAsRead(ctx echo.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockNotificationRepository) MarkAllAsRead(ctx echo.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func NewMockNotificationRepository() *MockNotificationRepository {
	return new(MockNotificationRepository)
}

func TestGetNotificationByID_Success(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	notificationId := "1"
	expectedNotification := &models.NotificationModel{
		Id:      "1",
		UserID:  "2",
		Title:   "Test Title",
		Message: "Test Message",
		IsRead:  false,
	}

	service := services.NewNotificationService(mockRepo)

	mockRepo.On("GetNotificationByID", mock.Anything, notificationId).Return(expectedNotification, nil)

	// Execution
	result, err := service.GetNotificationByID(nil, "1")

	// Assertion
	assert.Nil(t, err)
	assert.Equal(t, expectedNotification, result)
	mockRepo.AssertExpectations(t)
}

/*
func TestGetNotificationByID_NotFound(t *testing.T) {
	// Configura el mock del repositorio
	mockRepo := new(MockNotificationRepository)
	notificationId := "1"
	expectedErr := errors.New("Notification not found")

	// Configura el servicio con el mock del repositorio
	service := services.NewNotificationService(mockRepo)

	// Configura el mock del repositorio para el escenario de not found
	mockRepo.On("GetNotificationByID", mock.Anything, notificationId).Return(&models.NotificationModel{}, expectedErr)

	// Ejecución
	_, err := service.GetNotificationByID(nil, "1")

	// Verificación
	// assert.Nil(t, result)
	assert.Error(t, err)
}
*/

func TestGetNotificationsByUserID_Success(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	userId := "1"
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

	service := services.NewNotificationService(mockRepo)

	mockRepo.On("GetNotificationsByUserID", mock.Anything, userId).Return(expectedNotifications, nil)

	// Execution
	result, err := service.GetNotificationsByUserID(nil, "1")

	// Assertion
	assert.Nil(t, err)
	assert.Equal(t, expectedNotifications, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateNotification_Success(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	notification := &models.NotificationModel{
		Id:      "1",
		UserID:  "1",
		Title:   "Test Title",
		Message: "Test Message",
		IsRead:  false,
	}

	service := services.NewNotificationService(mockRepo)

	mockRepo.On("CreateNotification", mock.Anything, notification).Return(nil)

	// Execution
	err := service.CreateNotification(nil, notification)

	// Assertion
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMarkAsRead_Success(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	notificationId := "1"

	service := services.NewNotificationService(mockRepo)

	mockRepo.On("MarkAsRead", mock.Anything, notificationId).Return(nil)

	// Execution
	err := service.MarkAsRead(nil, "1")

	// Assertion
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMarkAllAsRead_Success(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	userId := "1"

	service := services.NewNotificationService(mockRepo)

	mockRepo.On("MarkAllAsRead", mock.Anything, userId).Return(nil)

	// Execution
	err := service.MarkAllAsRead(nil, "1")

	// Assertion
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
