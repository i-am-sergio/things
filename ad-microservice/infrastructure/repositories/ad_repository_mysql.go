package repositories

import (
	"ad-microservice/domain/models"
	"ad-microservice/domain/repository"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLConfig representa la configuración básica para conectarse a MySQL.
type MySQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string

	DB *gorm.DB // Agregar un campo para almacenar la conexión a la base de datos
}

// check implementation of interface
var _ repository.AdRepositoryInterface = &MySQLConfig{}

func SetMysql() *MySQLConfig {
	return &MySQLConfig{
		Host:     "roundhouse.proxy.rlwy.net",
		Port:     "58427",
		Username: "root",
		Password: "zblFuWrDprQNqiIffZpQIgkkTgofxSiF",
		Database: "railway",
	}
}

// ConnectDB inicializa y conecta a la base de datos MySQL utilizando la configuración proporcionada.
func (config *MySQLConfig) ConnectDB() error {
	// Crear el DSN (Data Source Name)
	DSN := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Conectar a la base de datos
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	config.DB = db // Guardar la instancia de la conexión en el struct

	println("Connected to database: ", db.Name())
	return nil
}

//Method interaction with db

// CreateAd crea un nuevo anuncio y lo guarda en la base de datos
func (config *MySQLConfig) CreateAd(newAd models.Add) error {
	if config.DB == nil { // Verificar si la conexión a la base de datos está establecida
		if err := config.ConnectDB(); err != nil {
			return err
		}
	}

	// Insertar en la base de datos
	if result := config.DB.Create(&newAd); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllAd obtiene todos los anuncios de la base de datos
func (config *MySQLConfig) GetAllAd() (*[]models.Add, error) {
	if config.DB == nil { // Verificar si la conexión a la base de datos está establecida
		if err := config.ConnectDB(); err != nil {
			return nil, err
		}
	}

	// Obtener todos los anuncios de la base de datos
	var adds []models.Add
	result := config.DB.Find(&adds)
	if result.Error != nil {
		return nil, result.Error // Devolver el error
	}

	return &adds, nil
}

// GetAddByIDProduct obtiene un anuncio por su ID de producto
func (config *MySQLConfig) GetAddByIDProduct(productID string) (*models.Add, error) {
	if config.DB == nil { // Verificar si la conexión a la base de datos está establecida
		if err := config.ConnectDB(); err != nil {
			return nil, err
		}
	}

	var add models.Add
	result := config.DB.Where("product_id = ?", productID).First(&add)
	if result.Error != nil {
		return nil, result.Error
	}
	return &add, nil
}

// UpdateAddData actualiza un anuncio con los nuevos datos si son diferentes de los existentes en la base de datos
func (config *MySQLConfig) UpdateAddData(updatedAdd models.Add) error {
	if config.DB == nil { // Verificar si la conexión a la base de datos está establecida
		if err := config.ConnectDB(); err != nil {
			return err
		}
	}

	// Actualizar el anuncio en la base de datos
	result := config.DB.Model(&models.Add{}).Where("product_id = ?", updatedAdd.ProductID).Updates(updatedAdd)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteAddByProductID elimina un anuncio por su ID
func (config *MySQLConfig) DeleteAddByProductID(productID string) error {
	if config.DB == nil { // Verificar si la conexión a la base de datos está establecida
		if err := config.ConnectDB(); err != nil {
			return err
		}
	}

	// Eliminar el Add con el product_id especificado
	result := config.DB.Where("product_id = ?", productID).Delete(&models.Add{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("no se encontró ningún registro con el product_id %s", productID)
		}
		return result.Error
	}

	return nil
}
