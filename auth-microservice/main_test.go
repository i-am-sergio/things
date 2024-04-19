package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeApp(t *testing.T) {

	e, port := Run()

	assert.NotNil(t, e)
	assert.Equal(t, "8001", port)
}
