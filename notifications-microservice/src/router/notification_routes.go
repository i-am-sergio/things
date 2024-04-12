package router

import (
	"notifications-microservice/src/controllers"
	"notifications-microservice/src/services"

	"github.com/labstack/echo/v4"
)

func NotificationRoutes(e *echo.Echo, notificationService services.NotificationService) {

	notifications := e.Group("/notifications")

	// Asignar el controlador al grupo de notificaciones
	notificationController := controllers.NewNotificationController(notificationService)

	// Routes
	notifications.GET("/:notification_id", notificationController.GetNotificationByID)
	notifications.GET("/user/:user_id", notificationController.GetNotificationsByUserID)
	notifications.POST("", notificationController.CreateNotification)
	notifications.PUT("/markAsRead/:notification_id", notificationController.MarkAsRead)
	notifications.PUT("/markAllAsRead/:user_id", notificationController.MarkAllAsRead)
}
