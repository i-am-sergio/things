package services

import "notifications-microservice/models"

type NotificationService interface {
	GetNotificationsByUserID(userID string) ([]models.Notification, error)
	GetNotificationByID(notificationID string) (models.Notification, error)
	MarkAsRead(notificationID string) error
	MarkAllAsRead(userID string) error
}
