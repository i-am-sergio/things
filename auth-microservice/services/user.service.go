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

	existingUser, err := GetUserByIdAuth(id)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return existingUser, nil
}

func ChangeUserRole(id string, newRole models.Role) (*models.User, error) {
	user, err := GetUserByIdAuth(id)
	if err != nil {
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

	return user, nil
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

func GetUserByIdAuth(idAuth string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("id_auth = ?", idAuth).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
