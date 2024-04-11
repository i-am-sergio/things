package services

import (
	"auth-microservice/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type RepositoryFunc interface {
	GetAllUsers() ([]models.User, int)
	CreateUser(user *models.User) (*models.User, int)
	UpdateUser(id string, updatedUser *models.User) (*models.User, int)
	ChangeUserRole(id string, newRole models.Role) (*models.User, int)
	GetUserByIdAuth(idAuth string) (*models.User, int)
}
type Database interface {
	First(dest interface{}, conds ...interface{}) error
	Save(value interface{}) error
	Create(value interface{}) error
	FindPreloaded(relation string, dest interface{}, conds ...interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	Delete(value interface{}) error
	DeleteByID(model interface{}, id interface{}) error
	Where(query interface{}, args ...interface{}) Database
	Model(value interface{}) Database
	Update(attrs ...interface{}) error
}

type DBClient struct {
	DB Database
}

func (d *DBClient) GetAllUsers() ([]models.User, int) {
	var users []models.User
	if err := d.DB.Find(&users); err != nil {
		// En caso de error, devuelve un código de estado HTTP interno del servidor
		return nil, http.StatusInternalServerError
	}
	// En caso de éxito, devuelve los usuarios y el código de estado HTTP OK
	return users, http.StatusOK
}
func (d *DBClient) CreateUser(user *models.User) (*models.User, int) {
	// Crea el usuario en la base de datos
	if err := d.DB.Create(user); err != nil {
		// En caso de error, devuelve el error y el código de estado HTTP interno del servidor
		return nil, http.StatusInternalServerError
	}

	// En caso de éxito, devuelve el usuario creado y el código de estado HTTP OK
	return user, http.StatusOK
}

func (d *DBClient) UpdateUser(id string, updatedUser *models.User) (*models.User, int) {
	// Buscar el usuario a actualizar
	var userToUpdate models.User
	if err := d.DB.Where("id_auth = ?", id).First(&userToUpdate); err != nil {
		// Si hay un error al buscar el usuario, devuelve el código de estado HTTP interno del servidor
		return nil, http.StatusInternalServerError
	}

	// Actualizar el usuario con los nuevos datos
	userToUpdate = *updatedUser
	// Actualiza otros campos según sea necesario

	// Guardar los cambios en la base de datos
	if err := d.DB.Save(&userToUpdate); err != nil {
		// Si hay un error al guardar los cambios, devuelve el código de estado HTTP interno del servidor
		return nil, http.StatusInternalServerError
	}

	return &userToUpdate, http.StatusOK
}

func (d *DBClient) GetUserByIdAuth(idAuth string) (*models.User, int) {
	var user models.User

	// Ejecutar la consulta utilizando GORM
	result := d.DB.Where("id_auth = ?", idAuth).First(&user)
	if result != nil {
		if errors.Is(result, gorm.ErrRecordNotFound) {
			// Devolver un código de estado HTTP de recurso no encontrado si no se encuentra el usuario
			return nil, http.StatusNotFound
		}
		// Devolver un código de estado HTTP interno del servidor si hay algún otro error
		return nil, http.StatusInternalServerError
	}

	// Devolver el usuario encontrado con el código de estado OK
	return &user, http.StatusOK
}

func (d *DBClient) ChangeUserRole(id string, newRole models.Role) (*models.User, int) {
	// Ejecutar la consulta utilizando GORM
	result := d.DB.Model(&models.User{}).Where("id_auth = ?", id).Update("role", newRole)
	if result != nil {
		// Devolver un código de estado HTTP interno del servidor si hay un error en la actualización
		return nil, http.StatusInternalServerError
	}

	// No hay error, lo que significa que la actualización fue exitosa
	// Devuelve el usuario actualizado con el código de estado OK
	return &models.User{IdAuth: id, Role: newRole}, http.StatusOK
}
