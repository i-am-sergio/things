package routes

import (
	"product-microservice/controllers"

	"github.com/labstack/echo/v4"
)

func CommentRoutes(e *echo.Echo) {
	comments := e.Group("/comments")
	comments.POST("", controllers.CreateComment)
	comments.GET("", controllers.GetComments)
	comments.GET("/:id", controllers.GetCommentByID)
	comments.GET("/products/:id", controllers.GetCommentsByProductID)
	comments.DELETE("/:id", controllers.DeleteComment)
	comments.PUT("/:id", controllers.UpdateComment)
}