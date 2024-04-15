package routes_test

import (
	"ad-microservice/app/controllers"
	"ad-microservice/app/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	e := echo.New()

	// Mock del controlador AdHandler
	mockController := &controllers.AdHandler{}

	// Configuración de las rutas con el mock del controlador
	routes.SetupRoutes(e, mockController)

	// Prueba de la ruta GET para obtener todos los anuncios
	req, err := http.NewRequest(http.MethodGet, "/ads", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Nil(t, err)
	// assert.Equal(t, http.StatusOK, rec.Code)
}

// func TestGetByIDProduct(t *testing.T) {
// 	e := echo.New()

// 	// Mock del controlador AdHandler
// 	mockController := &controllers.AdHandler{}

// 	// Configuración de las rutas con el mock del controlador
// 	routes.SetupRoutes(e, mockController)

// 	// Prueba de la ruta GET para obtener un anuncio por su ID
// 	req, err := http.NewRequest(http.MethodGet, "/ads/123", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	e.ServeHTTP(rec, req)
// 	assert.Nil(t, err)
// }
// func TestGetVerify(t *testing.T) {
// 	e := echo.New()

// 	// Mock del controlador AdHandler
// 	mockController := &controllers.AdHandler{}

// 	// Configuración de las rutas con el mock del controlador
// 	routes.SetupRoutes(e, mockController)

// 	// Prueba de la ruta GET para obtener todos los anuncios verificados
// 	req, err := http.NewRequest(http.MethodGet, "/ads/verify-requests", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	e.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusOK, rec.Code)

// }
// func TestPutAd(t *testing.T) {
// 	e := echo.New()

// 	// Mock del controlador AdHandler
// 	mockController := &controllers.AdHandler{}

// 	// Configuración de las rutas con el mock del controlador
// 	routes.SetupRoutes(e, mockController)

// 	// Prueba de la ruta PUT para actualizar un anuncio por su ID
// 	req, err := http.NewRequest(http.MethodPut, "/ads/123", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	e.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusOK, rec.Code)

// }
// func TestPostCreatAd(t *testing.T) {
// 	e := echo.New()

// 	// Mock del controlador AdHandler
// 	mockController := &controllers.AdHandler{}

// 	// Configuración de las rutas con el mock del controlador
// 	routes.SetupRoutes(e, mockController)

// 	// Prueba de la ruta POST
// 	req, err := http.NewRequest(http.MethodPost, "/ads", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	e.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusOK, rec.Code)

// }
// func TestDeleteAd(t *testing.T) {
// 	e := echo.New()

// 	// Mock del controlador AdHandler
// 	mockController := &controllers.AdHandler{}

// 	// Configuración de las rutas con el mock del controlador
// 	routes.SetupRoutes(e, mockController)

// 	// Prueba de la ruta DELETE para eliminar un anuncio por su ID
// 	req, err := http.NewRequest(http.MethodDelete, "/ads/123", nil)
// 	assert.NoError(t, err)
// 	rec := httptest.NewRecorder()
// 	e.ServeHTTP(rec, req)
// 	assert.Equal(t, http.StatusOK, rec.Code)

// }
