package routes

import (
	"product-microservice/controllers"

	"github.com/labstack/echo/v4"
)

func CommentRoutes(e *echo.Echo) {
	comments := e.Group("/comments")
	comments.POST("", controllers.CreateComment)
	comments.GET("/:id", controllers.GetCommentsByProductID)
	comments.DELETE("/:id", controllers.DeleteComment)
	comments.PUT("/rate/:id", controllers.UpdateProductRating)
	comments.PUT("/:id", controllers.UpdateComment)
}