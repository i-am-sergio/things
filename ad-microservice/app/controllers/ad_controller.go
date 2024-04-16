package controllers

import (
	"ad-microservice/app/services"
	"ad-microservice/domain/models"

	// "ad-microservice/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AdHandler maneja las solicitudes relacionadas con los ADS
type AdHandler struct {
	adService services.AdServiceInterface
}

// NewStudentHandler crea una nueva instancia de AdHandler
func NewAdHandler(adService services.AdServiceInterface) *AdHandler {
	return &AdHandler{
		adService: adService,
	}
}

var errorMessage string = "Internal Server Error"

// var notFoundMessage string = "Parameter Invalid"
func (s *AdHandler) CreateAdd(c echo.Context) error {
	// Crear una instancia de models.Add
	ad := new(models.Add)
	// Deserializar los datos del cuerpo de la solicitud en la instancia de models.Add
	if err := c.Bind(ad); err != nil {
		// Si hay un error al procesar los datos, devolver un código de estado HTTP 400 (Bad Request)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "no se pudieron procesar los datos"})
	}
	// Llamar a CreateAdService con el valor de models.Add
	if err := s.adService.CreateAdService(*ad); err != nil {
		// Si hay un error al guardar en la base de datos, devolver un código de estado HTTP 500 (Internal Server Error)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "no se pudo guardar en la base de datos"})
	}
	// Si no hay errores, devolver un código de estado HTTP 201 (Created)
	return c.JSON(http.StatusCreated, "Creado exitosamente")
}

// GetAddByIDProduct obtiene un anuncio por su ID de producto
func (s *AdHandler) GetAddByIdProduct(c echo.Context) error {
	// Obtener el valor de productId
	productID := c.Param("id")

	// Obtener el anuncio del servicio
	add, err := s.adService.GetAddByIDProductService(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener el anuncio"})
	}

	// Devolver el anuncio en formato JSON
	return c.JSON(http.StatusOK, add)
}

func (s *AdHandler) GetAllAdds(c echo.Context) error {
	// Declarar una variable para almacenar el resultado de GetAllAdService
	adds, err := s.adService.GetAllAdService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, adds)
}

// UpdateAddData actualiza un anuncio con los nuevos datos
func (s *AdHandler) UpdateAddData(c echo.Context) error {
	// Obtener el valor de productId
	productID := c.Param("id")

	// Leer el cuerpo de la solicitud para obtener los nuevos datos
	var updatedAdd models.Add
	if err := c.Bind(&updatedAdd); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request body"})
	}

	// Actualizar el anuncio con los nuevos datos
	if err := s.adService.UpdateAddDataService(productID, updatedAdd); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el anuncio"})
	}

	return c.NoContent(http.StatusAccepted)
}

// DeleteAddByID elimina un anuncio por su ID
func (s *AdHandler) DeleteAddByID(c echo.Context) error {
	// Obtener el valor de productId
	productID := c.Param("id")

	// Eliminar el anuncio por su ID
	if err := s.adService.DeleteAddByIDProductService(productID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al eliminar el anuncio"})
	}

	return c.NoContent(http.StatusAccepted)
}

// func GetAllAddsToShow(c echo.Context) error {
// 	adds, err := service.SelectPremiumProducts()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
// 	}
// 	return c.JSON(http.StatusOK, adds)
// }
