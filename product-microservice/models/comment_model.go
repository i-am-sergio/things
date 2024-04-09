package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" validate:"required" json:"id"`
	UserID    uint    `gorm:"index:idx_user_product" validate:"required" json:"user_id"`
	ProductID uint    `gorm:"index:idx_user_product" validate:"required" json:"product_id"`
	Comment   string  `validate:"required,min=1,max=1000" json:"comment"`
	Rating    float64 `validate:"required,gte=0,lte=5" json:"rating"`
}

func (c *Comment) TableName() string {
	return "comments"
}