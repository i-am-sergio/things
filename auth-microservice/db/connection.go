package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	// Obtener valores de las variables de entorno
	hostname := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Configurar la cadena de conexión DNS
	dns := "host=" + hostname + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port

	// Conectar a la base de datos
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected")
}
