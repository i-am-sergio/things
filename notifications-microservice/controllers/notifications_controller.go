package controllers

import (
	"net/http"
	"notifications-microservice/models"
	"notifications-microservice/services"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateNotification crea una nueva notificación
func CreateNotification(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Decodificar el cuerpo de la solicitud para obtener los datos de la notificación
		var notification models.Notification
		if err := c.Bind(&notification); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error al decodificar la solicitud"})
		}
		// Call the service to create a new notification
		if err := services.CreateNotificationService(notification, client); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al crear la notificación"})
		}
		// Responder con un código de estado 201 (Created) si la notificación se creó correctamente
		return c.JSON(http.StatusCreated, echo.Map{"message": "Notificación creada exitosamente"})
	}
}

// GetNotificationByID devuelve detalles específicos de una notificación
func GetNotificationByID(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		notificationID := c.Param("notification_id")
		notification, err := services.GetNotificationByIDService(notificationID, client)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Notification not found"})
		}
		return c.JSON(http.StatusOK, notification)
	}
}

// GetNotificationsByUserID devuelve todas las notificaciones de un usuario específico
func GetNotificationsByUserID(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		notifications, err := services.GetNotificationsByUserIDService(userID, client)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get notifications"})
		}
		return c.JSON(http.StatusOK, notifications)
	}
}

// MarkAsRead marca una notificación como leída
func MarkAsRead(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		notificationID := c.Param("notification_id")
		if err := services.MarkAsReadService(notificationID, client); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to mark notification as read"})
		}
		return c.NoContent(http.StatusOK)
	}
}

// MarkAllAsRead marca todas las notificaciones como leídas para un usuario específico
func MarkAllAsRead(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("user_id")
		if err := services.MarkAllAsReadService(userID, client); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark all notifications as read"})
		}
		return c.NoContent(http.StatusOK)
	}
}
