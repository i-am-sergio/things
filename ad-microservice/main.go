package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"ad-microservice/app/controllers"
	"ad-microservice/app/routes"
	"ad-microservice/domain/models"
	"ad-microservice/domain/service"
	"ad-microservice/infrastructure/repositories"
)

func main() {

	// Inicializar la configuración de la base de datos MySQL
	mysqlConfig := repositories.MySqlDB()

	// Conectar a la base de datos
	if err := mysqlConfig.ConnectDB(); err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
	}

	// AutoMigrar el modelo Add para crear la tabla en la base de datos
	if err := mysqlConfig.DB.AutoMigrate(&models.Add{}); err != nil {
		fmt.Println("Error al migrar el modelo Add:", err)
	}

	// Inicializar el repositorio de anuncios de MySQL
	var adRepo repositories.MySQLConfig = *mysqlConfig // Desreferenciar el puntero

	// Inicializar el servicio de anuncios
	adService := service.NewAdService(adRepo)

	// Inicializar el manejador de anuncios
	adHandler := controllers.NewAdHandler(adService)

	// Inicializar Echo
	e := echo.New()

	// Middleware
	//  e.Use(middleware.Logger())
	//  e.Use(middleware.Recover())

	// Ruta raíz
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Adds!")
	})
	// Configurar rutas
	routes.SetupRoutes(e, adHandler)

	// Iniciar el servidor
	fmt.Println("Servidor escuchando en el puerto 8003...")
	e.Logger.Fatal(e.Start(":8003"))

	// db.Init()

	// db.DB.AutoMigrate(&models.Add{})
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, Adds!")
	// })
	// routes.CommentRoutes(e)

	// e.Logger.Fatal(e.Start(":8003"))
}
