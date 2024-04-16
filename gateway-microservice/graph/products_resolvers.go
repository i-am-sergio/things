package graph

import (
	"context"
	"encoding/json"
	"gateway-microservice/graph/model"
	"net/http"
)

func (r *Resolver) GetProducts(ctx context.Context) ([]*model.Product, error) {
	// Realizar una solicitud HTTP GET a la API localhost:8005/products
	resp, err := http.Get("http://productsmcsv:8002/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta JSON en un modelo de product
	var products []*model.Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}