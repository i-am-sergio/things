package repository

import "gorm.io/gorm"

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