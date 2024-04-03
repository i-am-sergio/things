package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint	`gorm:"not null" validate:"required"`
	UserID    uint	`gorm:"not null" validate:"required"`
	ProductID uint	`gorm:"index;not null" validate:"required"`
	Comment   string	`gorm:"not null" validate:"required"`
	Rating    float64	`gorm:"not null" validate:"required"`
}