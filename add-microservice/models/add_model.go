package models

import (
	"time"

	"gorm.io/gorm"
)

type Add struct {
	gorm.Model
	ProductID int       `json:"productId"`
	Price     float64   `json:"price"`
	Time      int       `json:"time"`
	Date      time.Time `json:"date"`
	UserID    int       `json:"userId"`
	View      int       `json:"view"`
}
