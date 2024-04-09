package services

import (
	"context"
	"log"
	"notifications-microservice/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetNotificationByIDService devuelve detalles específicos de una notificación.
func GetNotificationByIDService(notificationID string, client *mongo.Client) (*models.Notification, error) {
	// Convertir el ID de string a tipo ObjectID de MongoDB
	objectID, err := primitive.ObjectIDFromHex(notificationID)
	if err != nil {
		log.Println("Invalid notification ID:", err)
		return nil, err
	}

	// Obtener la colección de notificaciones
	coll := client.Database("things").Collection("notifications")

	// Definir el contexto y realizar la búsqueda
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Buscar la notificación por ID
	var notification models.Notification
	err = coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&notification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Notification not found!")
			return nil, err // La notificación no fue encontrada
		}
		log.Println("Error al buscar la notificación:", err)
		return nil, err
	}

	return &notification, nil
}
