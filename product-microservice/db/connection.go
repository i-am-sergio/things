package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

type DBInterface interface {
	AutoMigrate(dst ...interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Save(value interface{}) error
	Create(value interface{}) error
	FindPreloaded(relation string, dest interface{}, conds ...interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	FindWithCondition(dest interface{}, query string, args ...interface{}) error
	Delete(value interface{}) error
	DeleteWithCondition(model interface{}, query string, args ...interface{}) error
	DeleteByID(model interface{}, id interface{}) error
}

type GormDBClient struct {
	DB *gorm.DB
	EnvLoader EnvLoader
	Connector  DBConnector
}

func (g *GormDBClient) Init(envLoader EnvLoader) *gorm.DB {
	g.EnvLoader = envLoader
	err := g.EnvLoader.LoadEnv()
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %v", err))
	}
	DSN := os.Getenv("DB_DSN")
	db, err := g.Connector.Open(DSN)
	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Connected to database.")
	}
	g.DB = db
	return db
}

func (g *GormDBClient) AutoMigrate(models ...interface{}) error {
	return g.DB.AutoMigrate(models...)
}

func (g *GormDBClient) First(dest interface{}, conds ...interface{}) error {
	return g.DB.First(dest, conds...).Error
}

func (g *GormDBClient) Save(value interface{}) error {
	return g.DB.Save(value).Error
}

func (g *GormDBClient) Create(value interface{}) error {
	return g.DB.Create(value).Error
}

func (g *GormDBClient) FindPreloaded(relation string, dest interface{}, conds ...interface{}) error {
	return g.DB.Preload(relation).Find(dest, conds...).Error
}

func (g *GormDBClient) Find(dest interface{}, conds ...interface{}) error {
	return g.DB.Find(dest, conds...).Error
}

func (g *GormDBClient) FindWithCondition(dest interface{}, query string, args ...interface{}) error {
	return g.DB.Where(query, args...).Find(dest).Error
}

func (g *GormDBClient) Delete(value interface{}) error {
	return g.DB.Delete(value).Error
}

func (g *GormDBClient) DeleteWithCondition(model interface{}, query string, args ...interface{}) error {
	return g.DB.Where(query, args...).Delete(model).Error
}

func (g *GormDBClient) DeleteByID(model interface{}, id interface{}) error {
	return g.DB.Delete(model, id).Error
}

var Client DBInterface

func Init(envLoader EnvLoader) {
	gormClient := &GormDBClient{
        EnvLoader: envLoader,
        Connector: &GormConnector{},
    }
    Client = gormClient
    gormClient.Init(envLoader)
}