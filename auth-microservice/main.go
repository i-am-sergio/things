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

// // DBConnectorImplementation es una implementación concreta de la interfaz DBConnector.
// type DBConnectorImplementation struct{}

// // DBConnection conecta a la base de datos utilizando el DNS proporcionado.
// func (d *DBConnectorImplementation) DBConnection(dns string) (*gorm.DB, error) {
// 	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
// 	if err != nil {
// 		log.Println("Failed to connect to database:", err)
// 		return nil, err
// 	}

// 	log.Println("DB connected")
// 	return db, nil
// }

func main() {
	e, port := Run()
	e.Logger.Fatal(e.Start(":" + port))
}

func Run() (*echo.Echo, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}
	if err := utils.Init(".env"); err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}
	dns := os.Getenv("DB_DNS")
	// connector := &DBConnectorImplementation{}
	conn, err := db.DBConnection(dns)
	conn.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	conn.Model(&models.User{}).Create(&models.User{
		Name: "Pepito", IdAuth: "1", Email: "turroncito_de_azucar@mock.com", Password: "brrr", Image: "https://res.cloudinary.com/dhocrtxvp/image/upload/v1712604420/users/20240408142659_flower.jpg", Ubication: "Contigo", Role: "ADMIN",
	})
	conn.Model(&models.User{}).Create(&models.User{
		Name: "Kiko", IdAuth: "2", Email: "testing@mock.com", Password: "comida", Image: "https://res.cloudinary.com/dhocrtxvp/image/upload/v1712459607/users/20240406221326_Captura%20desde%202023-12-09%2016-07-41.png", Ubication: "Sin ti", Role: "ADMIN",
	})
	repo := repository.NewUserRepository(conn)
	userService := services.NewUserService(repo)
	userController := controllers.NewUserController(userService)
	e := echo.New()

	routes.UsersRoutes(e, userController)
	port := os.Getenv("PORT")
	return e, port
}
