package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// Mocking the MongoDB URI for testing
	testURI := "mongodb://mongo:VkUwUwbMCWFpxsVbgPAZAlCtkpQnHXCq@roundhouse.proxy.rlwy.net:39308"

	client := ConnectDB(testURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if client is not nil
	assert.NotNil(t, client, "Client should not be nil")

	// Check if we can ping the MongoDB server
	err := client.Ping(ctx, nil)
	assert.NoError(t, err, "Ping to MongoDB failed")
}
