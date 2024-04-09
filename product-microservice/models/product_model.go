package models

import (
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    ID          uint    `gorm:"primaryKey;unique" json:"ID"`
	UserID      uint    `gorm:"index" json:"user_id" validate:"required"`
    State       bool    `gorm:"default:true" json:"state" `
    Status      bool    `gorm:"default:false" json:"status" validate:"required"`
    Name        string  `json:"name" validate:"required,min=5,max=100"`
    Description string  `json:"description" validate:"required,min=10,max=1000"`
    Image       string  `json:"image" validate:"required"`
    Category    string  `json:"category" validate:"required"`
    Price       float64 `json:"price" validate:"required,min=0"`
	Rate        float64 `gorm:"default:0.0" json:"rate" validate:"required,gte=0,lte=5"`
    Ubication   string  `json:"ubication" validate:"required,min=1"`
    Comments    []Comment `gorm:"foreignKey:ProductID" json:"comments" validate:"dive"`
}

func (p *Product) TableName() string {
    return "products"
}