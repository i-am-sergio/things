package db

import (
	"fmt"
	"os"
	"product-microservice/repository"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Client repository.DBInterface

type DBConnector interface {
    Open(dsn string) (*gorm.DB, error)
}
type GormConnector struct{}
func (g *GormConnector) Open(dsn string) (*gorm.DB, error) {
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

type EnvLoader interface {
    LoadEnv() error
}
type DotEnvLoader struct{}
func (d *DotEnvLoader) LoadEnv() error {
    return godotenv.Load()
}

type ConnectionRepository struct {
    EnvLoader   EnvLoader
    DBConnector DBConnector
}
func NewConnectionRepository(envLoader EnvLoader, dbConnector DBConnector) *ConnectionRepository {
    return &ConnectionRepository{
        EnvLoader:   envLoader,
        DBConnector: dbConnector,
    }
}

func (cr *ConnectionRepository) Init() (*gorm.DB, error) {
	if err := cr.EnvLoader.LoadEnv(); err != nil {
        return nil, fmt.Errorf("error loading .env file: %v", err)
    }
	DSN := os.Getenv("DB_DSN")
	db, err := cr.DBConnector.Open(DSN)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }
	fmt.Println("Connected to database.")
    return db, nil
}

func Init(envLoader EnvLoader, dbConnector DBConnector) error {
	repo := NewConnectionRepository(envLoader, dbConnector)
    db, err := repo.Init()
    if err != nil {
        return fmt.Errorf("failed to initialize database: %v", err)
    }
	gormClient := &repository.GormDBClient{
        DB: db,	
    }
    Client = gormClient
	return nil
}