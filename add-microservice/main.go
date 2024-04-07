package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"add-microservice/db"
	"add-microservice/models"
	"add-microservice/routes"
)

func main() {
	db.Init()

	db.DB.AutoMigrate(&models.Add{})
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Proucts!")
	})
	routes.CommentRoutes(e)

	e.Logger.Fatal(e.Start(":8002"))
}
