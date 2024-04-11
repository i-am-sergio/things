package service

import (
	"time"

	"ad-microservice/db"
	"ad-microservice/models"
)

// CreateAdService crea un nuevo anuncio y lo guarda en la base de datos
func CreateAdService(newAd models.Add) error {

	// Establecer la hora de creación
	now := time.Now()

	// Establecer la fecha y hora de creación y actualización
	newAd.CreatedAt = now
	newAd.UpdatedAt = now

	// Insertar en la base de datos
	if result := db.DB.Create(&newAd); result.Error != nil {
		return result.Error
	}

	return nil
}
