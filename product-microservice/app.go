package main

import (
	"fmt"
	"product-microservice/controllers"
	"product-microservice/db"
	"product-microservice/models"
	"product-microservice/routes"
	"product-microservice/services"

	"github.com/labstack/echo/v4"
)

type App struct {
	DB          db.RealDBInit
	Cloudinary  db.CloudinaryClient
	HTTPHandler *echo.Echo
}

func (a *App) Initialize() error {
	sql, err := a.DB.Init(&db.DotEnvLoader{}, &db.GormConnector{})
	if err != nil {
		return fmt.Errorf("failed to initialize DataBase: %w", err)
	}
	if err := sql.AutoMigrate(&models.Product{}, &models.Comment{}); err != nil {
		return fmt.Errorf("error AutoMigrate: %w", err)
	}
	if err := a.Cloudinary.InitCloudinary(&db.DotEnvLoader{}); err != nil {
		return fmt.Errorf("failed to initialize Cloudinary: %w", err)
	}
	productService := services.NewProductService(sql, a.Cloudinary)
	commentService := services.NewCommentService(sql)
	productController := controllers.NewProductController(productService)
	commentController := controllers.NewCommentController(commentService)
	routes.ProductRoutes(a.HTTPHandler, productController)
	routes.CommentRoutes(a.HTTPHandler, commentController)
	return nil
}

func (a *App) Run(port string) error {
	return a.HTTPHandler.Start(port)
}