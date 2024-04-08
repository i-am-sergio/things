package services

import (
	"auth-microservice/db"
	"auth-microservice/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) (*models.User, int) {

	if err := db.DB.Create(user).Error; err != nil {
		return nil, http.StatusInternalServerError
	}
	return user, http.StatusOK
}

func UpdateUser(id string, updatedUser *models.User) (*models.User, int) {

	existingUser, err := GetUserByIdAuth(id)
	if err != http.StatusOK {
		return nil, http.StatusInternalServerError
	}

	existingUser.Name = func() string {
		if updatedUser.Name != "" {
			return updatedUser.Name
		}
		return existingUser.Name
	}()

	existingUser.Email = func() string {
		if updatedUser.Email != "" {
			return updatedUser.Email
		}
		return existingUser.Email
	}()

	existingUser.Password = func() string {
		if updatedUser.Password != "" {
			return updatedUser.Password
		}
		return existingUser.Password
	}()

	existingUser.Image = func() string {
		if updatedUser.Image != "" {
			return updatedUser.Image
		}
		return existingUser.Image
	}()

	existingUser.Ubication = func() string {
		if updatedUser.Ubication != "" {
			return updatedUser.Ubication
		}
		return existingUser.Ubication
	}()

	if err := db.DB.Save(&existingUser).Error; err != nil {
		return nil, http.StatusInternalServerError
	}

	return existingUser, http.StatusOK
}

func ChangeUserRole(id string, newRole models.Role) (*models.User, int) {
	user, err := GetUserByIdAuth(id)
	if err != http.StatusOK {
		return nil, http.StatusInternalServerError
	}

	switch newRole {
	case models.RoleAdmin, models.RoleUser, models.RoleEnterprise:
	default:
		return nil, http.StatusBadRequest
	}
	user.Role = newRole
	if err := db.DB.Save(&user).Error; err != nil {
		return nil, http.StatusInternalServerError
	}

	return user, http.StatusOK
}
func GetUser(id string) (*models.User, int) {
	var user models.User

	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, http.StatusInternalServerError
	}
	if user.ID == 0 {
		return nil, http.StatusNotFound
	}
	return &user, http.StatusOK
}

func GetUserByIdAuth(idAuth string) (*models.User, int) {
	var user models.User
	if err := db.DB.Where("id_auth = ?", idAuth).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No se encontró ningún usuario con el ID de autenticación proporcionado, devolver error 404
			return nil, http.StatusNotFound
		}
		// Ocurrió un error diferente, devolver error 500
		return nil, http.StatusInternalServerError
	}
	// Usuario encontrado, devolver el usuario y el código de estado 200
	return &user, http.StatusOK
}
