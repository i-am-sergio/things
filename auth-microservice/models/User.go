package models

import "log"

type User struct {
	Name      string `gorm:"not null"`
	IdAuth    string `gorm:"not null;unique"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Image     string
	Ubication string
	Role      Role `gorm:"not null;default:USER"`
}

func (n *User) TestFunction() {
	log.Println("This is a test function")
}
