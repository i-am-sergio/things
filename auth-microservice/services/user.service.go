package services

import (
	"auth-microservice/db"
	"auth-microservice/models"
	"errors"
)

func CreateUser(user *models.User) (*models.User, error) {
	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id string, updatedUser *models.User) (*models.User, error) {
	var existingUser models.User
	if err := db.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Name = updatedUser.Name
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password
	existingUser.Image = updatedUser.Image
	existingUser.Ubication = updatedUser.Ubication

	if err := db.DB.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}
func ChangeUserRole(id string, newRole models.Role) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	switch newRole {
	case models.RoleAdmin, models.RoleUser, models.RoleEnterprise:
	default:
		return nil, errors.New("ROL NO VALIDO")
	}

	user.Role = newRole
	if err := db.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
func GetUser(id string) (*models.User, error) {
	var user models.User

	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, nil
	}
	return &user, nil
}
