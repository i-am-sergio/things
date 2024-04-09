package controllers

import (
	"add-microservice/db"
	"add-microservice/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var errorMessage string = "Internal Server Error"

// var notFoundMessage string = "Parameter Invalid"

func CreateAdd(c echo.Context) error {

	product := new(models.Add)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no se pudo procesar los datos"})
	}

	if result := db.DB.Create(&product); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no se guardo en la bd"})
	}

	return c.JSON(http.StatusCreated, product)

}

func GetAllAdds(c echo.Context) error {

	var adds []models.Add

	result := db.DB.Find(&adds)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, adds)
}

func GetAddByIdProduct(c echo.Context) error {

	// Obtener el valor de productId
	id := c.Param("id")
	fmt.Println("id:", id)
	var add models.Add
	result := db.DB.Where("product_id = ?", id).First(&add)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}

	return c.JSON(http.StatusOK, add)
}

func DeleteAddByID(c echo.Context) error {
	id := c.Param("id")
	var add models.Add

	// Obtener el Add con el id especificado
	result := db.DB.Where("product_id = ?", id).First(&add)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}

	// Eliminar el Add obtenido
	result = db.DB.Delete(&add)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}

	return c.NoContent(http.StatusAccepted)
}

func UpdateAddData(c echo.Context) error {
	// Obtener el productId de la URL
	productId := c.Param("id")

	// Buscar el anuncio en la base de datos por productId
	var add models.Add
	if result := db.DB.Where("product_id = ?", productId).First(&add); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Add not found"})
	}

	// Leer el cuerpo de la solicitud para obtener los nuevos datos
	var addUpdated models.Add
	if err := c.Bind(&addUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request body"})
	}

	// Actualizar los campos necesarios
	if addUpdated.Price != 0 {
		add.Price = addUpdated.Price
	}
	if addUpdated.Time != 0 {
		add.Time = addUpdated.Time
	}
	if !addUpdated.Date.IsZero() {
		add.Date = addUpdated.Date
	}

	// Actualizar el anuncio en la base de datos
	if result := db.DB.Save(&add); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update add"})
	}

	return c.NoContent(http.StatusAccepted)
}
