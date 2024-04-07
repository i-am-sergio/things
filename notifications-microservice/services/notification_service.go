package services

import (
	"context"
	"log"
	"notifications-microservice/db"
	"notifications-microservice/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetNotificationsByUserIDService devuelve todas las notificaciones de un usuario específico.
func GetNotificationsByUserIDService(userID string) ([]models.Notification, error) {
	// Obtener la conexión a la base de datos
	client := db.ConnectDB("your-mongo-uri")
	defer client.Disconnect(context.Background())

	// Obtener la colección de notificaciones
	coll := client.Database("things").Collection("notifications")

	// Definir el contexto y realizar la consulta
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Consultar las notificaciones por UserID
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

// GetNotificationByIDService devuelve detalles específicos de una notificación.
func GetNotificationByIDService(notificationID string) (*models.Notification, error) {
	// Obtener la conexión a la base de datos
	client := db.ConnectDB("your-mongo-uri")
	defer client.Disconnect(context.Background())

	// Obtener la colección de notificaciones
	coll := client.Database("things").Collection("notifications")

	// Definir el contexto y realizar la consulta
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Consultar la notificación por su ID
	var notification models.Notification
	err := coll.FindOne(ctx, bson.M{"_id": notificationID}).Decode(&notification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Notificación no encontrada")
			return nil, err // La notificación no fue encontrada
		}
		log.Println("Error al buscar notificación por ID:", err)
		return nil, err
	}

	return &notification, nil
}

// MarkAsReadService marca una notificación como leída.
func MarkAsReadService(notificationID string) error {
	// Obtener la conexión a la base de datos
	client := db.ConnectDB("your-mongo-uri")
	defer client.Disconnect(context.Background())

	// Obtener la colección de notificaciones
	coll := client.Database("things").Collection("notifications")

	// Definir el contexto y realizar la actualización
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Crear filtro para buscar la notificación por su ID
	filter := bson.M{"_id": notificationID}

	// Crear actualización para marcar la notificación como leída
	update := bson.M{"$set": bson.M{"isRead": true}}

	// Realizar la actualización en la base de datos
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error al marcar la notificación como leída:", err)
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

// MarkAllAsReadService marca todas las notificaciones como leídas para un usuario específico.
func MarkAllAsReadService(userID string) error {
	// Obtener la conexión a la base de datos
	client := db.ConnectDB("your-mongo-uri")
	defer client.Disconnect(context.Background())

	// Obtener la colección de notificaciones
	coll := client.Database("things").Collection("notifications")

	// Definir el contexto y realizar la actualización
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Crear filtro para buscar las notificaciones del usuario no leídas
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
		// Si no se ha modificado ninguna notificación, puede significar que no se encontraron notificaciones sin leer para ese usuario
		log.Println("No se encontraron notificaciones sin leer para el usuario proporcionado")
		return mongo.ErrNoDocuments
	}

	log.Println("Todas las notificaciones para el usuario", userID, "marcadas como leídas correctamente")
	return nil
}
