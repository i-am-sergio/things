package services

import (
	"notifications-microservice/src/models"
	"notifications-microservice/src/repositories"

	"github.com/labstack/echo/v4"
)

type NotificationService interface {
	GetNotificationByIDService(ctx echo.Context, id string) (*models.NotificationModel, error)
	GetNotificationsByUserIDService(ctx echo.Context, id string) ([]models.NotificationModel, error)
	CreateNotificationService(ctx echo.Context, notification *models.NotificationModel) error
	MarkAsReadService(ctx echo.Context, id string) error
	MarkAllAsReadService(ctx echo.Context, id string) error
}

// Class ---
type NotificationServiceImpl struct {
	repo repositories.NotificationRepository
}

// Constructor ---
func NewNotificationService(repo repositories.NotificationRepository) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		repo: repo,
	}
}

// Methods ---
func (s *NotificationServiceImpl) GetNotificationByIDService(ctx echo.Context, id string) (*models.NotificationModel, error) {
	return s.repo.GetNotificationByID(ctx, id)
}

func (s *NotificationServiceImpl) GetNotificationsByUserIDService(ctx echo.Context, userId string) ([]models.NotificationModel, error) {
	return s.repo.GetNotificationsByUserID(ctx, userId)
}

func (s *NotificationServiceImpl) CreateNotificationService(ctx echo.Context, notification *models.NotificationModel) error {
	return s.repo.CreateNotification(ctx, notification)
}

func (s *NotificationServiceImpl) MarkAsReadService(ctx echo.Context, id string) error {
	return s.repo.MarkAsRead(ctx, id)
}

func (s *NotificationServiceImpl) MarkAllAsReadService(ctx echo.Context, id string) error {
	return s.repo.MarkAllAsRead(ctx, id)
}
