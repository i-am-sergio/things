package service

import (
	"ad-microservice/db"
	"ad-microservice/models"
)

// GetAllAdService obtiene todos los anuncios de la base de datos
func GetAllAdService(adds *[]models.Add) error {
	// Obtener todos los anuncios de la base de datos
	result := db.DB.Find(&adds)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAddByIDProduct obtiene un anuncio por su ID de producto
func GetAddByIDProduct(productID string) (*models.Add, error) {
	var add models.Add
	result := db.DB.Where("product_id = ?", productID).First(&add)
	if result.Error != nil {
		return nil, result.Error
	}
	return &add, nil
}
