package service_test

import (
	"ad-microservice/domain/models"
	"ad-microservice/domain/service"
	"ad-microservice/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define una interfaz para el manejo de fechas.
type Clock interface {
	Now() time.Time
}

// Clock en tiempo real que implementa la interfaz Clock.
type RealClock struct{}

func (c *RealClock) Now() time.Time {
	return time.Now()
}

// MockClock para pruebas que implementa la interfaz Clock.
type MockClock struct {
	NowFunc func() time.Time
}

func (m *MockClock) Now() time.Time {
	if m.NowFunc != nil {
		return m.NowFunc()
	}
	return time.Time{}
}

func TestCreateAdService(t *testing.T) {
	// Crear una instancia del repositorio mock generado
	mockRepo := new(mocks.AdRepository)

	// Crear una instancia de MockClock con una función que devuelve una fecha específica.
	clock := &MockClock{
		NowFunc: func() time.Time {
			return time.Date(2024, time.April, 16, 4, 40, 11, 701289262, time.UTC)
		},
	}

	// Crear una instancia del servicio con el repositorio mock y el reloj falso
	serviceAd := service.NewAdService(mockRepo)

	// Crear un nuevo anuncio de prueba
	newAd := models.Add{
		ProductID: 123,
		Price:     99.99,
		Time:      60,
		Date:      clock.Now().AddDate(0, 0, 1),
		UserID:    456,
		View:      100,
	}

	// Establecer el comportamiento esperado del repositorio mock
	mockRepo.On("CreateAd", mock.AnythingOfType("models.Add")).Return(nil)

	// Llamar al método del servicio que estamos probando
	err := serviceAd.CreateAdService(newAd)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se llamó al método del mock como se esperaba
	mockRepo.AssertExpectations(t)

	// Verificar que la fecha del anuncio esté dentro de un margen de tiempo
	createdAt := clock.Now()
	timeDiff := createdAt.Sub(newAd.Date)
	assert.LessOrEqual(t, timeDiff.Seconds(), float64(10), "La diferencia de tiempo entre la fecha actual y la fecha del anuncio es mayor que 10 segundos")
}

func TestGetAddByIDProductService(t *testing.T) {
	// Crear una instancia del repositorio mock generado
	mockRepo := new(mocks.AdRepository)

	// Crear una instancia de MockClock con una función que devuelve una fecha específica.
	clock := &MockClock{
		NowFunc: func() time.Time {
			return time.Date(2024, time.April, 16, 4, 40, 11, 701289262, time.UTC)
		},
	}

	// Crear una instancia del servicio con el repositorio mock y el reloj falso
	serviceAd := service.NewAdService(mockRepo)

	// ID del producto de prueba
	productID := "123"

	// Crear un anuncio de prueba para el mock
	mockAdd := &models.Add{
		ProductID: 123,
		Price:     99.99,
		Time:      60,
		Date:      clock.Now(),
		UserID:    456,
		View:      100,
	}

	// Establecer el comportamiento esperado del repositorio mock
	mockRepo.On("GetAddByIDProduct", productID).Return(mockAdd, nil)

	// Llamar al método del servicio que estamos probando
	add, err := serviceAd.GetAddByIDProductService(productID)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se devuelva el anuncio correcto
	assert.Equal(t, mockAdd, add)

	// Verificar que se llamó al método del mock como se esperaba
	mockRepo.AssertExpectations(t)

	// Verificar que la fecha del anuncio esté dentro de un margen de tiempo
	createdAt := clock.Now()
	timeDiff := createdAt.Sub(mockAdd.Date)
	assert.LessOrEqual(t, timeDiff.Seconds(), float64(10), "La diferencia de tiempo entre la fecha actual y la fecha del anuncio es mayor que 10 segundos")
}

func TestGetAllAdService(t *testing.T) {
	// Crear una instancia del repositorio mock generado
	mockRepo := new(mocks.AdRepository)

	// Crear una instancia de MockClock con una función que devuelve una fecha específica.
	clock := &MockClock{
		NowFunc: func() time.Time {
			return time.Date(2024, time.April, 16, 4, 40, 11, 701289262, time.UTC)
		},
	}

	// Crear una instancia del servicio con el repositorio mock y el reloj falso
	serviceAd := service.NewAdService(mockRepo)

	// Crear una lista de anuncios de prueba para el mock
	mockAds := &[]models.Add{
		{
			ProductID: 123,
			Price:     99.99,
			Time:      60,
			Date:      clock.Now(),
			UserID:    456,
			View:      100,
		},
		{
			ProductID: 456,
			Price:     199.99,
			Time:      120,
			Date:      clock.Now().AddDate(0, 0, 1),
			UserID:    789,
			View:      200,
		},
	}

	// Establecer el comportamiento esperado del repositorio mock
	mockRepo.On("GetAllAd").Return(mockAds, nil)

	// Llamar al método del servicio que estamos probando
	ads, err := serviceAd.GetAllAdService()

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se devuelvan los anuncios correctos
	assert.Equal(t, mockAds, ads)

	// Verificar que se llamó al método del mock como se esperaba
	mockRepo.AssertExpectations(t)

	// Verificar que las fechas de los anuncios estén dentro de un margen de tiempo
	createdAt := clock.Now()
	for _, ad := range *mockAds {
		timeDiff := createdAt.Sub(ad.Date)
		assert.LessOrEqual(t, timeDiff.Seconds(), float64(10), "La diferencia de tiempo entre la fecha actual y la fecha del anuncio es mayor que 10 segundos")
	}
}

func TestUpdateAddDataService(t *testing.T) {
	// Crear una instancia del repositorio mock generado
	mockRepo := new(mocks.AdRepository)

	// Crear una instancia de MockClock con una función que devuelve una fecha específica.
	clock := &MockClock{
		NowFunc: func() time.Time {
			return time.Date(2024, time.April, 16, 4, 40, 11, 701289262, time.UTC)
		},
	}

	// Crear una instancia del servicio con el repositorio mock y el reloj falso
	serviceAd := service.NewAdService(mockRepo)

	// ID del producto de prueba
	productID := "123"

	// Anuncio actualizado de prueba
	updatedAd := models.Add{
		Price: 129.99,
		Time:  90,
		Date:  clock.Now().AddDate(0, 0, 2),
	}

	// Anuncio de prueba para el mock
	mockAdd := &models.Add{
		ProductID: 123,
		Price:     99.99,
		Time:      60,
		Date:      clock.Now(),
		UserID:    456,
		View:      100,
	}

	// Establecer el comportamiento esperado del repositorio mock
	mockRepo.On("GetAddByIDProduct", productID).Return(mockAdd, nil)
	mockRepo.On("UpdateAddData", mock.AnythingOfType("models.Add")).Return(nil)

	// Llamar al método del servicio que estamos probando
	err := serviceAd.UpdateAddDataService(productID, updatedAd)

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se llamó al método del mock como se esperaba
	mockRepo.AssertExpectations(t)

	// Verificar que la fecha del anuncio esté dentro de un margen de tiempo
	createdAt := clock.Now()
	timeDiff := createdAt.Sub(updatedAd.Date)
	assert.LessOrEqual(t, timeDiff.Seconds(), float64(10), "La diferencia de tiempo entre la fecha actual y la fecha del anuncio es mayor que 10 segundos")
}
