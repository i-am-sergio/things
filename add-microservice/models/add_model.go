package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Add struct {
	gorm.Model
	ProductID int       `json:"productId" validate:"required"`
	Price     float64   `json:"price" validate:"required,min=0"`
	Time      int       `json:"time" validate:"required,min=0"`
	Date      time.Time `json:"date" validate:"required"`
	UserID    int       `json:"userId" validate:"required"`
	View      int       `json:"view" validate:"required,min=0"`
}

// Función para validar el modelo Add
func (a *Add) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
