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
	// GIVEN
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

	// WHEN
	result, err := service.GetNotificationByIDService(nil, "1")

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, expectedNotification, result)
	mockRepo.AssertExpectations(t)
}

func TestGetNotificationsByUserID_Success(t *testing.T) {
	// GIVEN
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

	// WHEN
	result, err := service.GetNotificationsByUserIDService(nil, "1")

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, expectedNotifications, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateNotification_Success(t *testing.T) {
	// GIVEN
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

	// WHEN
	err := service.CreateNotificationService(nil, notification)

	// THEN
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMarkAsRead_Success(t *testing.T) {
	// GIVEN
	mockRepo := new(MockNotificationRepository)
	notificationId := "1"

	service := services.NewNotificationService(mockRepo)
	mockRepo.On("MarkAsRead", mock.Anything, notificationId).Return(nil)

	// WHEN
	err := service.MarkAsReadService(nil, "1")

	// THEN
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMarkAllAsRead_Success(t *testing.T) {
	// GIVEN
	mockRepo := new(MockNotificationRepository)
	userId := "1"

	service := services.NewNotificationService(mockRepo)
	mockRepo.On("MarkAllAsRead", mock.Anything, userId).Return(nil)

	// WHEN
	err := service.MarkAllAsReadService(nil, "1")

	// THEN
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
