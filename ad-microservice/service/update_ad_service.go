package service

import (
	"ad-microservice/db"
	"ad-microservice/models"
	"time"
)

// UpdateAddData actualiza un anuncio con los nuevos datos si son diferentes de los existentes en la base de datos
func UpdateAddData(productID string, updatedAdd models.Add) error {
	// Buscar el anuncio en la base de datos por productId
	var add models.Add
	if result := db.DB.Where("product_id = ?", productID).First(&add); result.Error != nil {
		return result.Error
	}

	// Actualizar los campos necesarios solo si son diferentes
	updateFields := make(map[string]interface{})

	// Verificar si el campo Price ha cambiado y actualizar si es así
	if updatedAdd.Price != 0 && updatedAdd.Price != add.Price {
		updateFields["price"] = updatedAdd.Price
	}

	// Verificar si el campo Time ha cambiado y actualizar si es así
	if updatedAdd.Time != 0 && updatedAdd.Time != add.Time {
		updateFields["time"] = updatedAdd.Time
	}

	// Actualizar el campo Date si es distinto de cero y diferente al valor actual
	if !updatedAdd.Date.IsZero() && !updatedAdd.Date.Equal(add.Date) {
		updateFields["date"] = updatedAdd.Date
	}

	// Establecer la fecha de actualización
	updateFields["updated_at"] = time.Now()

	// Actualizar el anuncio en la base de datos
	if result := db.DB.Model(&add).Updates(updateFields); result.Error != nil {
		return result.Error
	}

	return nil
}
