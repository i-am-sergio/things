package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name      string `gorm:"not null"`
	IdAuth    string `gorm:"not null;unique"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Image     string
	Ubication string
	Role      Role `gorm:"not null;default:USER"`
}
