package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeApp(t *testing.T) {
	// Mockear las funciones de carga de secretas y conexión a la base de datos si es necesario

	e, port := Run()

	// Verificar que el enrutador y otros componentes se hayan configurado correctamente
	assert.NotNil(t, e)
	assert.Equal(t, "8001", port)
}
