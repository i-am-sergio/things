package routes

import (
	"product-microservice/controllers"
	"product-microservice/db"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, cloudinary *db.Cloudinary) {
	products := e.Group("/products")
	products.POST("", controllers.CreateProduct(cloudinary))
	products.GET("", controllers.GetProducts)
	products.GET("/:id", controllers.GetProductsById)
	products.GET("", controllers.GetProductsByCategory)
	products.GET("/search", controllers.SearchProducts)
	products.PUT("/:id", controllers.UpdateProduct(cloudinary))
	products.DELETE("/:id", controllers.DeleteProduct)
	products.PUT("/:id/premium", controllers.Premium)
	products.GET("/premium", controllers.GetProductsPremium)
}