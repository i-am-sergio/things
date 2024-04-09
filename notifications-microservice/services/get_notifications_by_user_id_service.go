package services

import (
	"context"
	"log"
	"notifications-microservice/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetNotificationsByUserIDService devuelve todas las notificaciones de un usuario específico.
func GetNotificationsByUserIDService(userID string, client *mongo.Client) ([]models.Notification, error) {
	// Get the notifications collection
	coll := client.Database("things").Collection("notifications")

	// Define the context and perform the find
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the notifications by user ID
	cursor, err := coll.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		log.Println("Error al buscar notificaciones:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterar sobre los resultados y almacenar las notificaciones en un slice
	var notifications []models.Notification
	for cursor.Next(ctx) {
		var notification models.Notification
		if err := cursor.Decode(&notification); err != nil {
			log.Println("Error al decodificar notificación:", err)
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error en el cursor:", err)
		return nil, err
	}

	return notifications, nil
}
