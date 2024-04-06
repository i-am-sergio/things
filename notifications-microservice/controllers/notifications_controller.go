package controllers

import (
	"net/http"
	"notifications-microservice/services"

	"github.com/labstack/echo/v4"
)

// NotificationController maneja las solicitudes relacionadas con las notificaciones
type NotificationController struct {
	Service services.NotificationService
}

// GetNotificationsByUserID devuelve todas las notificaciones de un usuario específico
func (nc *NotificationController) GetNotificationsByUserID(c echo.Context) error {
	userID := c.Param("user_id")
	notifications, err := nc.Service.GetNotificationsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get notifications"})
	}
	return c.JSON(http.StatusOK, notifications)
}

// GetNotificationByID devuelve detalles específicos de una notificación
func (nc *NotificationController) GetNotificationByID(c echo.Context) error {
	notificationID := c.Param("notification_id")
	notification, err := nc.Service.GetNotificationByID(notificationID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Notification not found"})
	}
	return c.JSON(http.StatusOK, notification)
}

// MarkAsRead marca una notificación como leída
func (nc *NotificationController) MarkAsRead(c echo.Context) error {
	notificationID := c.Param("notification_id")
	if err := nc.Service.MarkAsRead(notificationID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark notification as read"})
	}
	return c.NoContent(http.StatusOK)
}

// MarkAllAsRead marca todas las notificaciones como leídas para un usuario específico
func (nc *NotificationController) MarkAllAsRead(c echo.Context) error {
	userID := c.Param("user_id")
	if err := nc.Service.MarkAllAsRead(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark all notifications as read"})
	}
	return c.NoContent(http.StatusOK)
}
