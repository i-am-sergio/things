package db

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// // DatabaseImpl es una implementación de la interfaz Database que envuelve la lógica de la función DBConnection.
// type Database interface {
// 	DBConnection() (*gorm.DB, error)
// }
// type DatabaseImpl struct {
// 	DB Database
// }

// func NewConnection(db Database) *DatabaseImpl {
// 	return &DatabaseImpl{
// 		DB: db,
// 	}
// }

// func (d *DatabaseImpl) DBConnection() (*gorm.DB, error) {
// 	dns := os.Getenv("DB_DNS")
// 	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
// 	if err != nil {
// 		log.Println("Failed to connect to database:", err)
// 		return nil, errors.New("failed to connect to database")
// 	}

// 	log.Println("DB connected")
// 	return db, nil
// }

func DBConnection(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, errors.New("failed to connect to database")
	}

	log.Println("DB connected")
	return db, nil
}
