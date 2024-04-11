package routes

import (
	"ad-microservice/controllers"

	"github.com/labstack/echo/v4"
)

// https://railway.app/project/edf31bd0-b4eb-40b0-9017-c7443c1be3b0/service/459c9d04-1540-4f26-be96-63a40fb1a7b5/data?state=table&table=adds
func CommentRoutes(e *echo.Echo) {
	ad := e.Group("/ads")
	ad.POST("", controllers.CreateAdd)
	ad.GET("verify-requests", controllers.GetAllAdds)
	ad.GET("/:id", controllers.GetAddByIdProduct)
	ad.GET("/products/premium", controllers.GetAllAddsToShow)
	ad.PUT("/:id", controllers.UpdateAddData)
	ad.DELETE("/:id", controllers.DeleteAddByID)
}
