package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MarkAllAsReadService marca todas las notificaciones como leídas para un usuario específico.
func MarkAllAsReadService(userID string, client *mongo.Client) error {
	// Get the notifications collection
	coll := client.Database("things").Collection("notifications")

	// Define the context and perform the update
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter to mark all notifications as read
	filter := bson.M{"userId": userID, "isRead": false}

	// Crear actualización para marcar todas las notificaciones como leídas
	update := bson.M{"$set": bson.M{"isRead": true}}

	// Realizar la actualización en la base de datos
	result, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Println("Error al marcar todas las notificaciones como leídas:", err)
		return err
	}

	if result.ModifiedCount == 0 {
		// Si no se ha modificado ninguna notificación,
		// puede significar que no se encontraron notificaciones sin leer para ese usuario
		log.Println("No se encontraron notificaciones sin leer para el usuario proporcionado")
		return mongo.ErrNoDocuments
	}

	log.Println("Todas las notificaciones para el usuario", userID, "marcadas como leídas correctamente")
	return nil
}
