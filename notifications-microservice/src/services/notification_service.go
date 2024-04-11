package services

import (
	"notifications-microservice/src/models"
	"notifications-microservice/src/repositories"

	"github.com/labstack/echo/v4"
)

type NotificationService interface {
	GetNotificationByID(ctx echo.Context, id string) (*models.NotificationModel, error)
	GetNotificationsByUserID(ctx echo.Context, id string) ([]models.NotificationModel, error)
	CreateNotification(ctx echo.Context, notification *models.NotificationModel) error
	MarkAsRead(ctx echo.Context, id string) error
	MarkAllAsRead(ctx echo.Context, id string) error
}

type NotificationServiceImpl struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		repo: repo,
	}
}

// Methods ---
func (s *NotificationServiceImpl) GetNotificationByID(ctx echo.Context, id string) (*models.NotificationModel, error) {
	return s.repo.GetNotificationByID(ctx, id)
}

func (s *NotificationServiceImpl) GetNotificationsByUserID(ctx echo.Context, id string) ([]models.NotificationModel, error) {
	return s.repo.GetNotificationsByUserID(ctx, id)
}

func (s *NotificationServiceImpl) CreateNotification(ctx echo.Context, notification *models.NotificationModel) error {
	return s.repo.CreateNotification(ctx, notification)
}

func (s *NotificationServiceImpl) MarkAsRead(ctx echo.Context, id string) error {
	return s.repo.MarkAsRead(ctx, id)
}

func (s *NotificationServiceImpl) MarkAllAsRead(ctx echo.Context, id string) error {
	return s.repo.MarkAllAsRead(ctx, id)
}
