package main

import (
	"auth-microservice/db"
	"auth-microservice/models"
	"auth-microservice/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	db.DBConnection()
	db.DB.AutoMigrate(&models.User{})

	e := echo.New()

	// Middleware global
	// Middleware específico para la ruta
	e.PUT("/users/role/:id", routes.ChangeRoleHandler)

	// Rutas
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", routes.CreateUserHandler)
	e.PUT("/users/:id", routes.UpdateUserHandler)
	e.GET("/users/:id", routes.GetUserHandler)

	// Iniciar el servidor
	e.Logger.Fatal(e.Start(":8001"))
}
