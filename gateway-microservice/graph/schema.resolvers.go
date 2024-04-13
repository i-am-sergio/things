package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"gateway-microservice/graph/model"
)

// CreateNotification is the resolver for the createNotification field.
func (r *mutationResolver) CreateNotification(ctx context.Context, userID string, title string, message string, isRead bool) (*model.Notification, error) {
	// panic(fmt.Errorf("not implemented: CreateNotification - createNotification"))
	return r.Resolver.CreateNotification(ctx, userID, title, message, isRead)
}

// GetNotificationByID is the resolver for the getNotificationById field.
func (r *queryResolver) GetNotificationByID(ctx context.Context, id string) (*model.Notification, error) {
	return r.Resolver.GetNotificationById(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
