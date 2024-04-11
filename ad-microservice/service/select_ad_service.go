package service

import (
	"ad-microservice/db"
	"ad-microservice/models"
	"errors"
	"sort"
	"strconv"
	"time"
)

// Wrapper struct para el modelo Add
type AddWrapper struct {
	*models.Add
}

// Función para verificar si una publicidad está activa
func (a *AddWrapper) IsActive() bool {
	// Calcula la fecha de finalización de la publicidad sumando la duración a la fecha de inicio
	expiration := a.Date.Add(time.Duration(a.Time))
	// Verifica si la fecha actual está antes de la fecha de finalización
	return time.Now().Before(expiration)
}

func SelectPremiumProducts() ([]models.Add, error) {
	var adds []models.Add

	// Obtener todas las publicidades
	result := db.DB.Find(&adds)
	if result.Error != nil {
		return nil, result.Error
	}

	// Seleccionar solo las publicidades activas
	var activeAdds []models.Add

	for _, add := range adds {
		// Creamos un AddWrapper para poder utilizar su método IsActive
		wrapper := AddWrapper{&add}
		if wrapper.IsActive() {
			activeAdds = append(activeAdds, add)
		}
	}

	// Definimos una función de comparación personalizada para sort.Slice
	less := func(i, j int) bool {
		// Primero comparamos por tiempo de duración
		if activeAdds[i].Time != activeAdds[j].Time {
			return activeAdds[i].Time < activeAdds[j].Time
		}
		// Si el tiempo de duración es igual, entonces comparamos por vistas
		return activeAdds[i].View < activeAdds[j].View
	}

	// Ordenamos las publicidades activas utilizando la función de comparación personalizada
	sort.Slice(activeAdds, less)

	return activeAdds, nil
}

func UpdateViewsAndReturnIDs(activeAdds []models.Add, numAdsRequired int) ([]string, error) {
	// Verificar si hay suficientes anuncios activos disponibles
	if len(activeAdds) < numAdsRequired {
		return nil, errors.New("no hay suficientes anuncios activos disponibles")
	}

	// Almacenar los IDs de los anuncios requeridos
	var adIDs []string
	for i := 0; i < numAdsRequired; i++ {
		adIDs = append(adIDs, strconv.Itoa(activeAdds[i].ProductID)) // Convertir int a string
	}

	// Aquí es donde deberías llamar a la función UpdateAdView si ya está definida
	UpdateAdView(adIDs)

	return adIDs, nil
}

// UpdateAdView incrementa las vistas en uno para cada anuncio dado su ID
func UpdateAdView(productIDs []string) error {
	for _, productID := range productIDs {
		// Buscar el anuncio en la base de datos por productId
		var add models.Add
		if result := db.DB.Where("product_id = ?", productID).First(&add); result.Error != nil {
			return result.Error
		}

		// Incrementar el campo View en uno
		add.View++

		// Actualizar el anuncio en la base de datos
		if result := db.DB.Save(&add); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
