package services

import (
	"ad-microservice/domain/models"
)

type AdService interface {
	CreateAdService(newAd models.Add) error
	GetAddByIDProductService(productID string) (*models.Add, error)
	GetAllAdService() (*[]models.Add, error)
	UpdateAddDataService(idProduct string, updatedAdd models.Add) error
	DeleteAddByIDProductService(productID string) error
}
