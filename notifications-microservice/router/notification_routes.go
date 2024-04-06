package router

import (
	"notifications-microservice/controllers"

	"github.com/labstack/echo/v4"
)

func NotificationRoutes(e *echo.Echo) {

	notifications := e.Group("/notifications")

	notifications.GET("/:user_id", controllers GetNotificationsByUserID)
	notifications.GET("/:notification_id", GetNotificationByID)
	notifications.PUT("/:notification_id", MarkAsRead)
	notifications.PUT("/:user_id", MarkAllAsRead)
}
