package routes

import (
	"product-microservice/controllers"

	"github.com/labstack/echo/v4"
)

func CommentRoutes(e *echo.Echo, c controllers.CommentController) {
	comments := e.Group("/comments")
	comments.POST("", c.CreateComment)
	comments.GET("", c.GetComments)
	comments.GET("/:id", c.GetCommentByID)
	comments.GET("/products/:id", c.GetCommentsByProductID)
	comments.DELETE("/:id", c.DeleteComment)
	comments.PUT("/:id", c.UpdateComment)
}