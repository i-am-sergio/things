package repositories

import (
	"notifications-microservice/src/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository interface {
	GetNotificationByID(ctx echo.Context, id string) (*models.NotificationModel, error)
	GetNotificationsByUserID(ctx echo.Context, id string) ([]models.NotificationModel, error)
	CreateNotification(ctx echo.Context, notification *models.NotificationModel) error
	MarkAsRead(ctx echo.Context, id string) error
	MarkAllAsRead(ctx echo.Context, id string) error
}

type NotificationRepositoryImpl struct {
	collection *mongo.Collection
}

func NewNotificationRepository(db *mongo.Database) *NotificationRepositoryImpl {
	return &NotificationRepositoryImpl{
		collection: db.Collection("notifications"),
	}
}

func (r *NotificationRepositoryImpl) GetNotificationByID(ctx echo.Context, id string) (*models.NotificationModel, error) {

	var notification models.NotificationModel
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx.Request().Context(), filter).Decode(&notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *NotificationRepositoryImpl) GetNotificationsByUserID(ctx echo.Context, id string) ([]models.NotificationModel, error) {

	var notifications []models.NotificationModel
	filter := bson.M{"userID": id}
	cursor, err := r.collection.Find(ctx.Request().Context(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx.Request().Context())

	for cursor.Next(ctx.Request().Context()) {
		var notification models.NotificationModel
		err := cursor.Decode(&notification)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (r *NotificationRepositoryImpl) CreateNotification(ctx echo.Context, notification *models.NotificationModel) error {
	// Generar un ID único para el post
	notification.Id = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(ctx.Request().Context(), notification)
	if err != nil {
		return err
	}

	return nil
}

func (r *NotificationRepositoryImpl) MarkAsRead(ctx echo.Context, id string) error {

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isRead": true}}
	result, err := r.collection.UpdateOne(ctx.Request().Context(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *NotificationRepositoryImpl) MarkAllAsRead(ctx echo.Context, id string) error {

	filter := bson.M{"userID": id}
	update := bson.M{"$set": bson.M{"isRead": true}}
	_, err := r.collection.UpdateMany(ctx.Request().Context(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
