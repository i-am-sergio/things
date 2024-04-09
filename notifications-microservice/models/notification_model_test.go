package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotificationModel(t *testing.T) {
	t.Run("TestCreateNotification", testCreateNotification)
	t.Run("TestNotificationEmptyFields", testNotificationEmptyFields)
	t.Run("TestNotificationSerialization", testNotificationSerialization)

	var notification Notification
	notification.TestFunction()
}

func testCreateNotification(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	notification := Notification{
		UserID:    "123",
		Title:     "Test Notification",
		Message:   "This is a test notification message",
		Image:     "test.jpg",
		IsRead:    false,
		Type:      "info",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	assert.Equal(t, "123", notification.UserID)
	assert.Equal(t, "Test Notification", notification.Title)
	assert.Equal(t, "This is a test notification message", notification.Message)
	assert.Equal(t, "test.jpg", notification.Image)
	assert.False(t, notification.IsRead)
	assert.Equal(t, "info", notification.Type)
	assert.Equal(t, createdAt, notification.CreatedAt)
	assert.Equal(t, updatedAt, notification.UpdatedAt)
}

func testNotificationEmptyFields(t *testing.T) {
	notification := Notification{}

	assert.Empty(t, notification.UserID)
	assert.Empty(t, notification.Title)
	assert.Empty(t, notification.Message)
	assert.Empty(t, notification.Image)
	assert.False(t, notification.IsRead)
	assert.Empty(t, notification.Type)
	assert.Zero(t, notification.CreatedAt)
	assert.Zero(t, notification.UpdatedAt)
}

func testNotificationSerialization(t *testing.T) {

	notification := Notification{
		UserID:  "123",
		Title:   "Test Notification",
		Message: "This is a test notification message",
		Image:   "test.jpg",
		IsRead:  false,
		Type:    "info",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(notification)
	assert.NoError(t, err)

	// Deserialize from JSON
	var deserializedNotification Notification
	err = json.Unmarshal(jsonData, &deserializedNotification)
	assert.NoError(t, err)

	// Verify deserialized fields
	assert.Equal(t, notification.UserID, deserializedNotification.UserID)
	assert.Equal(t, notification.Title, deserializedNotification.Title)
	assert.Equal(t, notification.Message, deserializedNotification.Message)
	assert.Equal(t, notification.Image, deserializedNotification.Image)
	assert.Equal(t, notification.IsRead, deserializedNotification.IsRead)
	assert.Equal(t, notification.Type, deserializedNotification.Type)
}
