package services

import (
	"context"
	"log"
	"time"

	"notifications-microservice/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreateNotificationService
func CreateNotificationService(notification models.Notification, client *mongo.Client) error {

	// Get the notifications collection
	coll := client.Database("things").Collection("notifications")

	// Define the context and perform the insert
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set the creation and update time
	notification.CreatedAt = time.Now()
	notification.UpdatedAt = time.Now()

	// Insert to database
	_, err := coll.InsertOne(ctx, notification)
	if err != nil {
		log.Println("Error to create notification!", err)
		return err
	}

	log.Println("Notification created successfully!")
	return nil
}
