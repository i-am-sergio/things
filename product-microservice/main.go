package main

import (
	"log"
	"net/http"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	if err:=db.Init(&db.DotEnvLoader{},&db.GormConnector{}); err != nil {
		log.Fatalf("Error al iniciar la base de datos: %v", err)
	}
	if err := db.Client.AutoMigrate(&models.Product{}, &models.Comment{}); err != nil {
		log.Fatalf("Error al realizar la migración: %v", err)
	}
	if err := db.InitCloudinary(&db.DotEnvLoader{}); err != nil {
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
