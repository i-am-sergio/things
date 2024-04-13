package main

import (
	"auth-microservice/controllers"
	"auth-microservice/db"
	"auth-microservice/models"
	"auth-microservice/repository"
	"auth-microservice/routes"
	"auth-microservice/services"
	"auth-microservice/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	// Inicializar Cloudinary
	if err := utils.Init(); err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}
	dns := os.Getenv("DB_DNS")
	conn, err := db.DBConnection(dns)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}
	conn.AutoMigrate(&models.User{})

	repo := repository.NewUserRepository(conn)

	userService := services.NewUserService(repo)

	userController := controllers.NewUserController(userService)

	e := echo.New()

	routes.UsersRoutes(e, userController)

	// Iniciar el servidor
	e.Logger.Fatal(e.Start(":8001"))
}
