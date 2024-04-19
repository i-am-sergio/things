package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway-microservice/graph/model"
	"net/http"
)

func (r *Resolver) GetUserById(ctx context.Context, id_auth string) (*model.User, error) {
	// Realizar una solicitud HTTP GET a la API localhost:8005/notifications/{id}
	resp, err := http.Get(fmt.Sprintf("http://authmcsv:8001/users/%s", id_auth))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta JSON en un modelo de notification
	var notification model.User
	if err := json.NewDecoder(resp.Body).Decode(&notification); err != nil {
		return nil, err
	}

	return &notification, nil
}
