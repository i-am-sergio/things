package main

import (
	"auth-microservice/controllers"
	"auth-microservice/routes"
	"auth-microservice/services"
	"log"
	"net/http"
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
	if err := services.Init(); err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	// Establecer la conexión con la base de datos
	// db.DBConnection()
	// db.DB.AutoMigrate(&models.User{})

	// Crear una instancia de Echo
	e := echo.New()

	// Ruta de inicio simple
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Users!")
	})

	// Obtener las variables de entorno para la conexión a la base de datos
	dialect := "postgres"
	dsn := "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"

	// Crear una instancia del repositorio
	repo, err := services.NewRepository(dialect, dsn, 10, 100) // Establece el número máximo de conexiones
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	// Crear una instancia del controlador de usuario con el servicio de repositorio
	userController := controllers.NewUserController(repo)

	// Asociar las rutas de usuarios con el controlador de usuarios
	routes.UsersRoutes(e, userController)

	// Iniciar el servidor
	e.Logger.Fatal(e.Start(":8001"))
}
