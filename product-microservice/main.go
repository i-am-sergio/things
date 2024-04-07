package main

import (
	"log"
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/routes"
	"product-microservice/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.Product{}, &models.Comment{})
	if err := service.Init(); err != nil {
        log.Fatalf("Failed to initialize Cloudinary: %v", err)
    }
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Proucts!")
	})
	routes.ProductRoutes(e)
	routes.CommentRoutes(e)
	e.Logger.Fatal(e.Start(":8002"))
}
