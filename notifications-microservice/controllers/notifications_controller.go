package controllers

import (
	"net/http"
	"notifications-microservice/services"

	"github.com/labstack/echo/v4"
)

// GetNotificationsByUserID devuelve todas las notificaciones de un usuario específico
func GetNotificationsByUserID(c echo.Context) error {
	userID := c.Param("user_id")
	notifications, err := services.GetNotificationsByUserIDService(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get notifications"})
	}
	return c.JSON(http.StatusOK, notifications)
}

// GetNotificationByID devuelve detalles específicos de una notificación
func GetNotificationByID(c echo.Context) error {
	notificationID := c.Param("notification_id")
	notification, err := services.GetNotificationByIDService(notificationID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Notification not found"})
	}
	return c.JSON(http.StatusOK, notification)
}

// MarkAsRead marca una notificación como leída
func MarkAsRead(c echo.Context) error {
	notificationID := c.Param("notification_id")
	if err := services.MarkAsReadService(notificationID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark notification as read"})
	}
	return c.NoContent(http.StatusOK)
}

// MarkAllAsRead marca todas las notificaciones como leídas para un usuario específico
func MarkAllAsRead(c echo.Context) error {
	userID := c.Param("user_id")
	if err := services.MarkAllAsReadService(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark all notifications as read"})
	}
	return c.NoContent(http.StatusOK)
}
