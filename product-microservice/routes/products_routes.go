package routes

import (
	"product-microservice/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	products := e.Group("/products")
	products.POST("", controllers.CreateProduct)
	products.GET("", controllers.GetProducts)
	products.GET("/:id", controllers.GetProductsById)
	products.GET("", controllers.GetProductsByCategory)
	products.GET("/search", controllers.SearchProducts)
	products.PUT("/:id", controllers.UpdateProduct)
	products.DELETE("/:id", controllers.DeleteProduct)
	products.PUT("/:id/premium", controllers.Premium)

}