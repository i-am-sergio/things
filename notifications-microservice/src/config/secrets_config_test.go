package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSecrets(t *testing.T) {
	port, mongoURI, _ := LoadSecrets()
	assert.NotEmpty(t, port, "Port should not be empty")
	assert.NotEmpty(t, mongoURI, "MongoURI should not be empty")
}

func TestLoadSecretError(t *testing.T) {
	_, _, err := LoadSecrets()
	assert.Nil(t, err, "Error should be nil")
}
