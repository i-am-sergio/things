package main

import (
	"auth-microservice/db"
	"auth-microservice/models"
	"auth-microservice/routes"
	"auth-microservice/services"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}
	if err := services.Init(); err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}
	db.DBConnection()
	db.DB.AutoMigrate(&models.User{})

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Users!")
	})

	routes.UsersRoutes(e)

	// Iniciar el servidor
	e.Logger.Fatal(e.Start(":8001"))
}
