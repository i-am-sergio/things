package routes

import (
	"ad-microservice/app/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, controllers *controllers.AdHandler) {
	ad := e.Group("/ads")
	ad.POST("", controllers.CreateAdd)
	ad.GET("/verify-requests", controllers.GetAllAdds)
	ad.GET("/:id", controllers.GetAddByIdProduct)
	//ad.GET("/products/premium", controllers.GetAllAddsToShow)
	ad.PUT("/:id", controllers.UpdateAddData)
	ad.DELETE("/:id", controllers.DeleteAddByID)
}
