package main

import (
	"net/http"
	"notifications-microservice/config"
	"notifications-microservice/db"
	"notifications-microservice/router"

	"github.com/labstack/echo/v4"
)

func main() {

	port, mongoURI := config.LoadSecrets()

	client := db.ConnectDB(mongoURI)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.NotificationRoutes(e, client)

	e.Logger.Fatal(e.Start(":" + port))
}
