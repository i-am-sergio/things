package controllers

import (
	"ad-microservice/models"
	"ad-microservice/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var errorMessage string = "Internal Server Error"

// var notFoundMessage string = "Parameter Invalid"
func CreateAdd(c echo.Context) error {
	// Crear una instancia de models.Add
	ad := new(models.Add)

	// Deserializar los datos del cuerpo de la solicitud en la instancia de models.Add
	if err := c.Bind(ad); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no se pudo procesar los datos"})
	}

	// Llamar a CreateAdService con el valor de models.Add
	if err := service.CreateAdService(*ad); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no se guardo en la bd"})
	}

	return c.JSON(http.StatusCreated, "Creado exitosamente")
}

func GetAllAdds(c echo.Context) error {

	var adds []models.Add

	if err := service.GetAllAdService(&adds); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, adds)
}

// GetAddByIDProduct obtiene un anuncio por su ID de producto
func GetAddByIdProduct(c echo.Context) error {
	// Obtener el valor de productId
	productID := c.Param("id")

	// Obtener el anuncio del servicio
	add, err := service.GetAddByIDProduct(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener el anuncio"})
	}

	// Devolver el anuncio en formato JSON
	return c.JSON(http.StatusOK, add)
}

// DeleteAddByID elimina un anuncio por su ID
func DeleteAddByID(c echo.Context) error {
	// Obtener el valor de productId
	productID := c.Param("id")

	// Eliminar el anuncio por su ID
	if err := service.DeleteAddByID(productID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al eliminar el anuncio"})
	}

	return c.NoContent(http.StatusAccepted)
}

// UpdateAddData actualiza un anuncio con los nuevos datos
func UpdateAddData(c echo.Context) error {
	// Obtener el valor de productId
	productID := c.Param("id")

	// Leer el cuerpo de la solicitud para obtener los nuevos datos
	var updatedAdd models.Add
	if err := c.Bind(&updatedAdd); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request body"})
	}

	// Actualizar el anuncio con los nuevos datos
	if err := service.UpdateAddData(productID, updatedAdd); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el anuncio"})
	}

	return c.NoContent(http.StatusAccepted)
}

func GetAllAddsToShow(c echo.Context) error {
	adds, err := service.SelectPremiumProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, adds)
}
