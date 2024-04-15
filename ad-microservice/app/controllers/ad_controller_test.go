package controllers_test

import (
	"ad-microservice/app/controllers"
	"ad-microservice/domain/models"
	"ad-microservice/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestGetAddByIdProductHandler(t *testing.T) {
	// Crear una instancia del servicio mock generado
	mockService := new(mocks.AdService)

	// Crear un nuevo anuncio de prueba
	expectedAdd := &models.Add{
		ProductID: 123,
		Price:     99.99,
		Time:      60,
		Date:      time.Now().AddDate(0, 0, 1),
		UserID:    456,
		View:      100,
	}

	// Mockear el servicio para que devuelva el anuncio esperado
	mockService.On("GetAddByIDProductService", "123").Return(expectedAdd, nil)

	// Crear una instancia del controlador con el servicio mock
	controller := controllers.NewAdHandler(mockService)

	// Crear una nueva solicitud HTTP para enviar al controlador
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/123", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("123")

	// Llamar al método del controlador que estamos probando
	err := controller.GetAddByIdProduct(ctx)

	// Verificar que no haya errores
	assert.NoError(t, err)
	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verificar que la respuesta JSON coincide con el anuncio esperado
	var responseAdd models.Add
	err = json.Unmarshal(rec.Body.Bytes(), &responseAdd)
	assert.NoError(t, err)
	assert.Equal(t, *expectedAdd, responseAdd)

	// Verificar que se llamó al método del mock como se esperaba
	mockService.AssertExpectations(t)
}

func TestGetAllAddsHandler(t *testing.T) {
	// Crear una instancia del servicio mock generado
	mockService := new(mocks.AdService)

	// Crear una lista de anuncios de prueba
	expectedAdds := []models.Add{
		{ProductID: 1, Price: 10.0},
		{ProductID: 2, Price: 20.0},
	}

	// Mockear el servicio para que devuelva la lista de anuncios esperada
	mockService.On("GetAllAdService").Return(&expectedAdds, nil)

	// Crear una instancia del controlador con el servicio mock
	controller := controllers.NewAdHandler(mockService)

	// Crear una nueva solicitud HTTP para enviar al controlador
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	// Llamar al método del controlador que estamos probando
	err := controller.GetAllAdds(ctx)

	// Verificar que no haya errores
	assert.NoError(t, err)
	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verificar que la respuesta JSON coincide con los anuncios esperados
	var responseAdds []models.Add
	err = json.Unmarshal(rec.Body.Bytes(), &responseAdds)
	assert.NoError(t, err)
	assert.Equal(t, expectedAdds, responseAdds)

	// Verificar que se llamó al método del mock como se esperaba
	mockService.AssertExpectations(t)
}

func TestUpdateAddDataHandler(t *testing.T) {
	// Crear una instancia del servicio mock generado
	mockService := new(mocks.AdService)

	// Crear un anuncio de prueba
	updatedAdd := models.Add{ProductID: 123, Price: 99.99}

	// Mockear el servicio para que no devuelva ningún error
	mockService.On("UpdateAddDataService", "123", updatedAdd).Return(nil)

	// Crear una instancia del controlador con el servicio mock
	controller := controllers.NewAdHandler(mockService)

	// Crear una nueva solicitud HTTP para enviar al controlador
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/123", strings.NewReader(`{"productID":123,"price":99.99}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("123")

	// Llamar al método del controlador que estamos probando
	err := controller.UpdateAddData(ctx)

	// Verificar que no haya errores
	assert.NoError(t, err)
	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusAccepted, rec.Code)

	// Verificar que se llamó al método del mock como se esperaba
	mockService.AssertExpectations(t)
}

func TestDeleteAddByIDHandler(t *testing.T) {
	// Crear una instancia del servicio mock generado
	mockService := new(mocks.AdService)

	// Mockear el servicio para que no devuelva ningún error
	mockService.On("DeleteAddByIDProductService", "123").Return(nil)

	// Crear una instancia del controlador con el servicio mock
	controller := controllers.NewAdHandler(mockService)

	// Crear una nueva solicitud HTTP para enviar al controlador
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/123", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("123")

	// Llamar al método del controlador que estamos probando
	err := controller.DeleteAddByID(ctx)

	// Verificar que no haya errores
	assert.NoError(t, err)
	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusAccepted, rec.Code)

	// Verificar que se llamó al método del mock como se esperaba
	mockService.AssertExpectations(t)
}
