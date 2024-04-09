package router

import (
	"notifications-microservice/controllers"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NotificationRoutes(e *echo.Echo, client *mongo.Client) {

	notifications := e.Group("/notifications")

	notifications.POST("", controllers.CreateNotification(client))

	notifications.GET("/id/:notification_id", controllers.GetNotificationByID(client))
	notifications.GET("/userId/:user_id", controllers.GetNotificationsByUserID(client))
	notifications.PUT("/markAsRead/:notification_id", controllers.MarkAsRead(client))
	notifications.PUT("/markAllAsRead/:user_id", controllers.MarkAllAsRead(client))
}
