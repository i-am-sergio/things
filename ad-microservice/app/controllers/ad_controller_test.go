package controllers_test

import (
	"ad-microservice/app/controllers"
	"ad-microservice/app/services/mocks"
	"ad-microservice/domain/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateAdd_Succes(t *testing.T) {
	// Crear una instancia del servicio mock generado
	mockService := new(mocks.AdService)

	// Creamos un nuevo anuncio para enviar al controlador
	newAd := &models.Add{
		ProductID: 123,
		Price:     99.99,
		Time:      60,
		Date:      time.Now().AddDate(0, 0, 1), // Usar el mismo tiempo que se espera en el mock
		UserID:    456,
		View:      100,
	}

	// Convertir el anuncio a JSON
	notificationJSON, _ := json.Marshal(newAd)

	// Crear una instancia del controlador con el servicio mock
	controller := controllers.NewAdHandler(mockService)

	// Crear una nueva solicitud HTTP para enviar al controlador
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(notificationJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Configurar el comportamiento esperado del mock
	mockService.On("CreateAdService", *newAd).Return(nil)

	// Llamar al método del controlador que estamos probando
	err := controller.CreateAdd(ctx)

	// Verificar que no haya errores
	assert.NoError(t, err)
	// Verificar que la respuesta tenga el código de estado correcto
	assert.Equal(t, http.StatusCreated, rec.Code)
	// Verificar que se llamó al método del mock como se esperaba
	mockService.AssertExpectations(t)
}
