package main

import (
	"errors"
	"notifications-microservice/src/config"
	"notifications-microservice/src/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interfaces para los mocks
type configLoaderInterface interface {
	LoadSecrets() (string, string, error)
}

type dbConnectorInterface interface {
	ConnectDB(uri string) (*mongo.Client, error)
}

// Mock para simular un error al cargar secretos
type mockConfigLoaderError struct{}

func (m *mockConfigLoaderError) LoadSecrets() (string, string, error) {
	return "", "", errors.New("error loading secrets")
}

// Mock para simular un error al conectar a la base de datos
type mockDBConnectorError struct{}

func (m *mockDBConnectorError) ConnectDB(uri string) (*mongo.Client, error) {
	return nil, errors.New("error connecting to database")
}

// Mock para simular una conexión exitosa a la base de datos
type mockDBConnectorSuccess struct{}

func (m *mockDBConnectorSuccess) ConnectDB(uri string) (*mongo.Client, error) {
	return db.ConnectDB(uri)
}

// Mock para simular una carga exitosa de secretos
type mockConfigLoaderSuccess struct{}

func (m *mockConfigLoaderSuccess) LoadSecrets() (string, string, error) {
	return config.LoadSecrets()
}

// Variables globales para las interfaces de los mocks
var configLoader configLoaderInterface = &mockConfigLoaderError{}
var dbConnector dbConnectorInterface = &mockDBConnectorError{}

func TestInitializeAppErrorLoadingSecrets(t *testing.T) {
	// Mockear las funciones de carga de secretas y conexión a la base de datos si es necesario
	configLoader = &mockConfigLoaderError{}
	dbConnector = &mockDBConnectorSuccess{}

	e, port, errSecrets, _ := initializeApp(
		configLoader.LoadSecrets,
		dbConnector.ConnectDB,
	)

	// Verificar que se haya retornado un error al cargar secretos
	assert.Nil(t, e)
	assert.Equal(t, "", port)
	assert.NotNil(t, errSecrets)
}

func TestInitializeAppErrorConnectingToDB(t *testing.T) {
	// Mockear las funciones de carga de secretas y conexión a la base de datos si es necesario
	configLoader = &mockConfigLoaderSuccess{}
	dbConnector = &mockDBConnectorError{}

	e, port, _, errDB := initializeApp(
		configLoader.LoadSecrets,
		dbConnector.ConnectDB,
	)

	// Verificar que se haya retornado un error al conectar a la base de datos
	assert.Nil(t, e)
	assert.Equal(t, "", port)
	assert.NotNil(t, errDB)
}

func TestInitializeAppSuccess(t *testing.T) {
	// Mockear las funciones de carga de secretas y conexión a la base de datos si es necesario
	configLoader = &mockConfigLoaderSuccess{}
	dbConnector = &mockDBConnectorSuccess{}

	e, port, _, _ := initializeApp(
		configLoader.LoadSecrets,
		dbConnector.ConnectDB,
	)

	// Verificar que no se haya retornado un error
	assert.NotNil(t, e)
	assert.Equal(t, "8005", port)
}
