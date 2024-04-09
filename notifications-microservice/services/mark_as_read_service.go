package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MarkAsReadService marca una notificación como leída.
func MarkAsReadService(notificationID string, client *mongo.Client) error {
	// Convertir el ID de string a tipo ObjectID de MongoDB
	objectID, err := primitive.ObjectIDFromHex(notificationID)
	if err != nil {
		log.Println("Invalid notification ID:", err)
		return err
	}

	// Obtener la colección de notificaciones
	coll := client.Database("things").Collection("notifications")

	// Definir el contexto y realizar la actualización
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Crear filtro para buscar la notificación por su ID
	filter := bson.M{"_id": objectID}

	// Crear actualización para marcar la notificación como leída
	update := bson.M{"$set": bson.M{"isRead": true}}

	// Realizar la actualización en la base de datos
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error to mark the notification as readed", err)
		return err
	}

	if result.ModifiedCount == 0 {
		// Si no se ha modificado ninguna notificación, puede significar que no se encontró ninguna notificación con ese ID
		log.Println("No se encontró ninguna notificación con el ID proporcionado")
		return mongo.ErrNoDocuments
	}

	log.Println("Notificación marcada como leída correctamente")
	return nil
}
