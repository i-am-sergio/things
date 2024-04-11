package probandoo

import (
	"auth-microservice/models"
	"net/http"
)

type Database interface {
	Find(dest interface{}, conds ...interface{}) error
}
type DBClient struct {
	DB Database
}

type RepositoryFunc interface {
	GetAllUsers() ([]models.User, int)
}

func (d *DBClient) GetAllUsers() ([]models.User, int) {
	var users []models.User
	println("holaaa")
	if err := d.DB.Find(&users); err != nil {
		// En caso de error, devuelve un código de estado HTTP interno del servidor
		return nil, http.StatusInternalServerError
	}
	// En caso de éxito, devuelve los usuarios y el código de estado HTTP OK
	return users, http.StatusOK
}
