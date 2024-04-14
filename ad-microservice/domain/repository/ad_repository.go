package repository

import (
	"ad-microservice/domain/models"
)

type AdRepository interface {
	ConnectDB() error
	GetAddByIDProduct(productID string) (*models.Add, error)
	GetAllAd() (*[]models.Add, error)
	CreateAd(newAd models.Add) error
	DeleteAddByProductID(productID string) error
	UpdateAddData(updatedAdd models.Add) error
	//SelectPremiumProducts() ([]models.Add, error)
	// DisconnectDB()
}
