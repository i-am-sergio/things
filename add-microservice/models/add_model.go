package models

import (
	"time"

	"gorm.io/gorm"
)

type Add struct {
	gorm.Model
	ID        int       `json:"ID"`
	ProductID int       `json:"productID"`
	Price     float64   `json:"price"`
	Time      time.Time `json:"time"`
	Date      time.Time `json:"date"`
	UserID    int       `json:"userId"`
}
