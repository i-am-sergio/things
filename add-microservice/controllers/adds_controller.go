package controllers

import (
	"add-microservice/db"
	"add-microservice/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var errorMessage string = "Internal Server Error"

// var notFoundMessage string = "Parameter Invalid"

func CreateAdd(c echo.Context) error {

	product := new(models.Add)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}

	if result := db.DB.Create(&product); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
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

func GetAddByID(c echo.Context) error {
	id := c.Param("id")
	var add models.Add

	result := db.DB.First(&add, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}

	return c.JSON(http.StatusOK, add)
}

func DeleteAddByID(c echo.Context) error {
	id := c.Param("id")
	var add models.Add

	result := db.DB.Delete(&add, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.NoContent(http.StatusAccepted)
}

func UpdateAddData(c echo.Context) error {
	id := c.Param("id")
	var add models.Add

	// Convertir el parámetro de la URL `id` a un tipo adecuado (por ejemplo, int)
	addID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid 'id' parameter"})
	}

	// Buscar el registro en la base de datos por su ID
	result := db.DB.First(&add, addID)
	if result.Error != nil {
		// Si el registro no se encuentra, devolver un error 404 (Not Found)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Add not found"})
	}

	// Leer el cuerpo de la solicitud para obtener el nuevo dato del campo Price
	var updateData map[string]string
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request body"})
	}

	// Verificar si se proporciona un nuevo valor para el campo Price en la solicitud
	newPriceStr, ok := updateData["price"]
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing 'price' field in request body"})
	}

	// Convertir el valor de newPriceStr a float64
	newPrice, err := strconv.ParseFloat(newPriceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value for 'price' field"})
	}

	// Actualizar el campo Price con el nuevo valor
	add.Price = newPrice

	// Guardar los cambios en la base de datos
	result = db.DB.Save(&add)
	if result.Error != nil {
		// Si hay un error al guardar, devolver un error 500 (Internal Server Error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update add"})
	}

	return c.NoContent(http.StatusAccepted)
}
