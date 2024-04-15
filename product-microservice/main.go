package main

import (
	"fmt"
	"log"
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
	sql, err := a.DB.Init(&db.DotEnvLoader{},&db.GormConnector{})
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

func runApp(db *db.RealDBInitImpl, cloudinary *db.Cloudinary, e *echo.Echo) error {
    app := &App{
        DB:          db,
        Cloudinary:  cloudinary,
        HTTPHandler: e,
    }
    if err := app.Initialize(); err != nil {
        return fmt.Errorf("failed to initialize the application: %w", err)
    }
    if err := app.Run(":8002"); err != nil {
        return fmt.Errorf("failed to start the server: %w", err)
    }
    return nil
}

func main() {
	cloudinary := &db.Cloudinary{
        Uploader: &db.CloudinaryUploaderAdapter{},
        API:      &db.CloudinaryService{},
    }
    e := echo.New()
    db := &db.RealDBInitImpl{}

    if err := runApp(db, cloudinary, e); err != nil {
        log.Fatal(err)
    }
}