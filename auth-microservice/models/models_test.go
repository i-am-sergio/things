package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels(t *testing.T) {
	t.Run("TestCreateuser", testCreateuser)
	t.Run("TestUserEmptyFields", testuserEmptyFields)
	t.Run("TestUserSerialize", testuserSerialization)
	var user User
	user.TestFunction()
}

func testCreateuser(t *testing.T) {

	user := User{
		Name:      "pepe",
		IdAuth:    "123",
		Email:     "qweqwe",
		Password:  "asdasd",
		Image:     "image",
		Ubication: "local",
		Role:      "USER",
	}

	assert.Equal(t, "pepe", user.Name)
	assert.Equal(t, "123", user.IdAuth)
	assert.Equal(t, "qweqwe", user.Email)
	assert.Equal(t, "image", user.Image)
	assert.Equal(t, "local", user.Ubication)
	assert.Equal(t, "USER", string(user.Role))
}

func testuserEmptyFields(t *testing.T) {
	user := User{}

	assert.Empty(t, user.Name)
	assert.Empty(t, user.IdAuth)
	assert.Empty(t, user.Email)
	assert.Empty(t, user.Image)
	assert.Empty(t, user.Ubication)
	assert.Empty(t, user.Role)
}

func testuserSerialization(t *testing.T) {

	user := User{
		Name:      "pepe",
		IdAuth:    "123",
		Email:     "qweqwe",
		Password:  "asdasd",
		Image:     "image",
		Ubication: "local",
		Role:      "USER",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(user)
	assert.NoError(t, err)

	// Deserialize from JSON
	var userSerialized User
	err = json.Unmarshal(jsonData, &userSerialized)
	assert.NoError(t, err)

	// Verify deserialized fields
	assert.Equal(t, user.Name, userSerialized.Name)
	assert.Equal(t, user.Email, userSerialized.Email)
	assert.Equal(t, user.IdAuth, userSerialized.IdAuth)
	assert.Equal(t, user.Image, userSerialized.Image)
	assert.Equal(t, user.Ubication, userSerialized.Ubication)
	assert.Equal(t, user.Role, userSerialized.Role)
}
