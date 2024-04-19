package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gateway-microservice/graph/model"
	"net/http"
)

func (r *Resolver) GetNotificationById(ctx context.Context, id string) (*model.Notification, error) {
	// Realizar una solicitud HTTP GET a la API localhost:8005/notifications/{id}
	resp, err := http.Get(fmt.Sprintf("http://notificationsmcsv:8005/notifications/%s", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta JSON en un modelo de notification
	var notification model.Notification
	if err := json.NewDecoder(resp.Body).Decode(&notification); err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *Resolver) CreateNotification(ctx context.Context, userID string, title string, message string, isRead bool) (*model.Notification, error) {
	// Crear un modelo de notificación con los datos proporcionados
	notification := model.Notification{
		UserID:  &userID,
		Title:   &title,
		Message: &message,
		IsRead:  &isRead,
	}

	// Serializar el modelo de notificación en JSON
	body, err := json.Marshal(notification)
	if err != nil {
		return nil, err
	}

	// Realizar una solicitud HTTP POST a la API localhost:8005/notifications
	resp, err := http.Post("http://notificationsmcsv:8005/notifications", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta JSON en un modelo de notificación
	if err := json.NewDecoder(resp.Body).Decode(&notification); err != nil {
		return nil, err
	}

	return &notification, nil
}
