package graph

import (
	"context"
	"encoding/json"
	"gateway-microservice/graph/model"
	"net/http"
)

func (r *Resolver) GetComments(ctx context.Context) ([]*model.Comment, error) {
	resp, err := http.Get("http://productsmcsv:8002/comments")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var comments []*model.Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, err
	}
	return comments, nil
}