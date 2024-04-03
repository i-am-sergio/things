package main

import (
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.Product{}, &models.Comment{})
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Proucts!")
	})
	routes.ProductRoutes(e)
	routes.CommentRoutes(e)
	e.Logger.Fatal(e.Start(":8002"))
}
