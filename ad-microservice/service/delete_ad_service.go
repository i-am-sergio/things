package service

import (
	"ad-microservice/db"
	"ad-microservice/models"
)

// DeleteAddByID elimina un anuncio por su ID
func DeleteAddByID(productID string) error {
	var add models.Add

	// Obtener el Add con el id especificado
	result := db.DB.Where("product_id = ?", productID).First(&add)
	if result.Error != nil {
		return result.Error
	}

	// Eliminar el Add obtenido
	result = db.DB.Delete(&add)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
