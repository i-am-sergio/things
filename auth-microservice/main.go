package main

import (
	"auth-microservice/controllers"
	"auth-microservice/db"
	"auth-microservice/repository"
	"auth-microservice/routes"
	"auth-microservice/services"
	"auth-microservice/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConnectorImplementation es una implementación concreta de la interfaz DBConnector.
type DBConnectorImplementation struct{}

// DBConnection conecta a la base de datos utilizando el DNS proporcionado.
func (d *DBConnectorImplementation) DBConnection(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	log.Println("DB connected")
	return db, nil
}

func main() {
	e, port := Run()
	e.Logger.Fatal(e.Start(":" + port))
}

func Run() (*echo.Echo, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}
	if err := utils.Init(); err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}
	dns := os.Getenv("DB_DNS")
	connector := &DBConnectorImplementation{}
	conn, err := db.DBConnection(connector, dns)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repo := repository.NewUserRepository(conn)
	userService := services.NewUserService(repo)
	userController := controllers.NewUserController(userService)
	e := echo.New()

	routes.UsersRoutes(e, userController)
	port := os.Getenv("PORT")
	return e, port
}
