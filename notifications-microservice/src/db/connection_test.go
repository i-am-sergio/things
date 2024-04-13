package db

import (
	"context"
	"notifications-microservice/src/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// Mocking the MongoDB URI for testing
	_, mongoURI, _ := config.LoadSecrets()

	client, err := ConnectDB(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if client is not nil
	assert.NotNil(t, client, "Client should not be nil")
	// Check if there is no error
	assert.Nil(t, err, "Error should be nil")

	// ping the client
	err = client.Ping(ctx, nil)
	// Check if there is no error
	assert.Nil(t, err, "Error should be nil")
}

func TestConnectDBFailure(t *testing.T) {
	// Mocking the MongoDB URI for testing
	mongoURI := "mongodb://localhost:27017/invalidurl"
	// Injecting an invalid URI
	_, err := ConnectDB(mongoURI)
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if client is nil
	assert.Nil(t, err, "Client should be nil")
}
