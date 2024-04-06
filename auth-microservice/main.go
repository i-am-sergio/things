package main

import (
	"auth-microservice/db"
	"auth-microservice/middleware"
	"auth-microservice/models"
	"auth-microservice/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Aplica el middleware de validación de JWT
		handler := middleware.EnsureValidToken()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Set("user", c.Get("user"))
			next(c)
		}))
		handler.ServeHTTP(c.Response(), c.Request())
		// Este return es opcional, dependiendo de si quieres que la ejecución continúe después de este middleware o no.
		return next(c)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	db.DBConnection()
	db.DB.AutoMigrate(&models.User{})

	e := echo.New()

	// Middleware global
	// e.Use(JWTMiddleware)

	// Middleware específico para la ruta
	e.PUT("/users/role/:id", routes.ChangeRoleHandler, JWTMiddleware)

	// Rutas
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", routes.CreateUserHandler, JWTMiddleware)
	e.PUT("/users/:id", routes.UpdateUserHandler)
	e.GET("/users/:id", routes.GetUserHandler, JWTMiddleware)

	// Iniciar el servidor
	e.Logger.Fatal(e.Start(":8001"))
}
