package routes

import (
	"add-microservice/controllers"

	"github.com/labstack/echo/v4"
)

// https://railway.app/project/edf31bd0-b4eb-40b0-9017-c7443c1be3b0/service/459c9d04-1540-4f26-be96-63a40fb1a7b5/data?state=table&table=adds
func CommentRoutes(e *echo.Echo) {
	comments := e.Group("/adds")
	comments.POST("", controllers.CreateAdd)
	comments.GET("", controllers.GetAllAdds)
	comments.GET("/:id", controllers.GetAddByIdProduct)
	comments.PUT("/:id", controllers.UpdateAddData)
	comments.DELETE("/:id", controllers.DeleteAddByID)
}
