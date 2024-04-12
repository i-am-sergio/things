package repository

import (
	"auth-microservice/models"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, id string, updatedUser *models.User) (*models.User, error)
	ChangeUserRole(ctx context.Context, id string, newRole models.Role) (*models.User, error)
	GetUserByIdAuth(ctx context.Context, idAuth string) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (d *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := d.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := d.DB.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserRepositoryImpl) UpdateUser(ctx context.Context, id string, updatedUser *models.User) (*models.User, error) {
	if err := d.DB.WithContext(ctx).Model(&models.User{}).Where("id_auth = ?", id).Updates(updatedUser).Error; err != nil {
		return nil, err
	}

	// Devolver el usuario actualizado
	return updatedUser, nil
}

func (d *UserRepositoryImpl) GetUserByIdAuth(ctx context.Context, idAuth string) (*models.User, error) {
	var user *models.User

	err := d.DB.WithContext(ctx).Where("id_auth = ?", idAuth).First(&user).Error

	return user, err
}
func (d *UserRepositoryImpl) ChangeUserRole(ctx context.Context, id string, newRole models.Role) (*models.User, error) {
	var user models.User
	result := d.DB.WithContext(ctx).Model(&user).Where("id_auth = ?", id).Update("role", newRole)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
