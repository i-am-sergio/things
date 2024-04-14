package service

import (
	"errors"
	"time"

	"ad-microservice/app/services"
	"ad-microservice/domain/models"
	"ad-microservice/infrastructure/repositories"
)

type adService struct {
	adRepo repositories.MySQLConfig
	ads    map[string]models.Add
}

// constructor
func NewAdService(repo ...repositories.MySQLConfig) services.AdService {
	if len(repo) > 0 {
		return &adService{
			adRepo: repo[0],
		}
	}
	return &adService{
		ads: make(map[string]models.Add),
	}
}

// METHODS
// check implementation of interface
var _ services.AdService = &adService{}

func (s *adService) CreateAdService(newAd models.Add) error {
	// Establecer la hora de creación
	now := time.Now()

	// Establecer la fecha y hora de creación y actualización
	newAd.CreatedAt = now
	newAd.UpdatedAt = now

	// Insertar en la base de datos
	result := s.adRepo.CreateAd(newAd) // Asignar el resultado a result
	if result != nil {
		return result
	}

	return nil
}

func (s *adService) GetAddByIDProductService(productID string) (*models.Add, error) {

	// Llamar al método GetAddByIDProduct en el repositorio
	add, err := s.adRepo.GetAddByIDProduct(productID)
	if err != nil {
		return nil, err
	}
	return add, nil
}

func (s *adService) GetAllAdService() (*[]models.Add, error) {
	// Obtener todos los anuncios de la base de datos
	ads, err := s.adRepo.GetAllAd()
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func (s *adService) UpdateAddDataService(productID string, updatedAd models.Add) error {

	// Buscar el anuncio en la base de datos por productId
	add, err := s.adRepo.GetAddByIDProduct(productID)
	if err != nil {
		return err
	}

	// Copiar el anuncio existente en el objeto updatedFields y actualizar los campos necesarios
	updatedFields := *add

	// Verificar si el campo Price ha cambiado y actualizar si es así
	if updatedAd.Price != 0 && updatedAd.Price != updatedFields.Price {
		updatedFields.Price = updatedAd.Price
	}

	// Verificar si el campo Time ha cambiado y actualizar si es así
	if updatedAd.Time != 0 && updatedAd.Time != updatedFields.Time {
		updatedFields.Time = updatedAd.Time
	}

	// Verificar si el campo Date ha cambiado y actualizar si es así
	if !updatedAd.Date.IsZero() && !updatedAd.Date.Equal(updatedFields.Date) {
		updatedFields.Date = updatedAd.Date
	}

	// Establecer la fecha de actualización
	updatedFields.UpdatedAt = time.Now()

	// Actualizar el anuncio en la base de datos
	result := s.adRepo.UpdateAddData(updatedFields)
	if result != nil {
		return errors.New(result.Error()) // Convertir la cadena a un error
	}

	return nil
}

// // UpdateAdView incrementa las vistas en uno para cada anuncio dado su ID
// func (s *adService) UpdateAdView(productIDs []string) error {
// 	for _, productID := range productIDs {
// 		// Buscar el anuncio en la base de datos por productId
// 		var add models.Add
// 		if result := s.adRepo.GetAddByIDProduct(productID); result.Error != nil {
// 			return result.Error
// 		}

// 		// Incrementar el campo View en uno
// 		add.View++

// 		// Actualizar el anuncio en la base de datos
// 		if result := s.adRepo.UpdateAddData(add); result.Error != nil {
// 			return result.Error
// 		}
// 	}
// 	return nil
// }

func (s *adService) DeleteAddByIDProductService(productID string) error {
	// Obtener el Add con el id especificado
	result := s.adRepo.DeleteAddByProductID(productID)
	if result != nil {
		return errors.New(result.Error()) // Convertir la cadena a un error
	}

	return nil
}
