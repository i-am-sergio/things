package routes

import (
	"product-microservice/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, c controllers.ProductController) {
	products := e.Group("/products")
	products.POST("", c.CreateProduct)
	products.GET("", c.GetProducts)
	products.GET("/:id", c.GetProductsById)
	products.GET("", c.GetProductsByCategory)
	products.GET("/search", c.SearchProducts)
	products.PUT("/:id", c.UpdateProduct)
	products.DELETE("/:id", c.DeleteProduct)
	products.PUT("/:id/premium", c.Premium)
	products.GET("/premium", c.GetProductsPremium)
}