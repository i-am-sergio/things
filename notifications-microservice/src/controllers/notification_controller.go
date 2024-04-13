package controllers

import (
	"net/http"

	"notifications-microservice/src/models"
	"notifications-microservice/src/services"

	"github.com/labstack/echo/v4"
)

// Class
type NotificationController struct {
	service services.NotificationService
}

// Constructor ---
func NewNotificationController(service services.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

// Methods ---

func (nc *NotificationController) GetNotificationByID(c echo.Context) error {
	notificationID := c.Param("notification_id")
	notification, err := nc.service.GetNotificationByIDService(c, notificationID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Notification not found"})
	}
	return c.JSON(http.StatusOK, notification)
}

func (nc *NotificationController) GetNotificationsByUserID(c echo.Context) error {
	userID := c.Param("user_id")
	notifications, err := nc.service.GetNotificationsByUserIDService(c, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Failed to get notifications"})
	}
	return c.JSON(http.StatusOK, notifications)
}

func (nc *NotificationController) CreateNotification(c echo.Context) error {
	var notification models.NotificationModel
	if err := c.Bind(&notification); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to decode request"})
	}
	if err := nc.service.CreateNotificationService(c, &notification); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create notification"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Notification created successfully"})
}

func (nc *NotificationController) MarkAsRead(c echo.Context) error {
	notificationID := c.Param("notification_id")
	if err := nc.service.MarkAsReadService(c, notificationID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to mark notification as read"})
	}
	return c.NoContent(http.StatusOK)
}

func (nc *NotificationController) MarkAllAsRead(c echo.Context) error {
	userID := c.Param("user_id")
	if err := nc.service.MarkAllAsReadService(c, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to mark all notifications as read"})
	}
	return c.NoContent(http.StatusOK)
}
