package router

import (
	"notifications-microservice/src/controllers"

	"github.com/labstack/echo/v4"
)

// Class ---
type Router struct {
	controller *controllers.NotificationController
}

// Constructor ---
func NewRouter(controller *controllers.NotificationController) *Router {
	return &Router{controller: controller}
}

// Methods ---
func (r *Router) NotificationRoutes(e *echo.Echo) {

	// Group
	notifications := e.Group("/notifications")

	// Rutas
	notifications.GET("/:notification_id", r.controller.GetNotificationByID)
	notifications.GET("/user/:user_id", r.controller.GetNotificationsByUserID)
	notifications.POST("", r.controller.CreateNotification)
	notifications.PUT("/markAsRead/:notification_id", r.controller.MarkAsRead)
	notifications.PUT("/markAllAsRead/:user_id", r.controller.MarkAllAsRead)
}
