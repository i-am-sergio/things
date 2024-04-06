package main

import (
	"context"
	"log"
	"net/http"
	"notifications-microservice/db"
	"notifications-microservice/router"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")

	client := db.ConnectDB(mongoURI)

	coll := client.Database("things").Collection("notifications")

	coll.InsertOne(context.TODO(), map[string]string{"name": "test"})

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.NotificationRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
