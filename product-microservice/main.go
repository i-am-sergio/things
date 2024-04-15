package main

import (
	"log"
	"product-microservice/db"

	"github.com/labstack/echo/v4"
)

func main() {
    if err := RunServer(); err != nil {
        log.Fatal("failed to start the server: ", err)
    }
}

func RunServer() error {
    cloudinary := &db.Cloudinary{
        Uploader: &db.CloudinaryUploaderAdapter{},
        API:      &db.CloudinaryService{},
    }
    e := echo.New()
    db := &db.RealDBInitImpl{}
    app := &App{
        DB:          db,
        Cloudinary:  cloudinary,
        HTTPHandler: e,
    }
    if err := app.Initialize(); err != nil {
        return err
    }
    return app.Run(":8002")
}