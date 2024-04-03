package models

import (
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    ID          uint    `gorm:"not null;unique" validate:"required"`
	UserID      uint    `gorm:"not null" validate:"required"`
    Status      bool    `gorm:"default:true"`
    Name        string  `gorm:"not null" validate:"required"`
    Description string  `gorm:"not null" validate:"required"`
    Category    string  `gorm:"not null" validate:"required"`
    Price       float64 `gorm:"not null" validate:"required"`
	Ratings     float64 `gorm:"default:0.0"`
    Comments    []Comment `gorm:"foreignKey:ProductID"`
}