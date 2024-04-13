package db

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// DBConnector define la interfaz para la conexión a la base de datos.
type DBConnector interface {
	DBConnection(dns string) (*gorm.DB, error)
}

// DBConnection conecta a la base de datos utilizando el conector proporcionado.
func DBConnection(connector DBConnector, dns string) (*gorm.DB, error) {
	db, err := connector.DBConnection(dns)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, errors.New("failed to connect to database")
	}

	log.Println("DB connected")
	return db, nil
}
