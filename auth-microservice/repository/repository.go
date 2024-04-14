package repository

import (
	"auth-microservice/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(ctx echo.Context) ([]models.User, error)
	CreateUser(ctx echo.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx echo.Context, id string, updatedUser *models.User) (*models.User, error)
	ChangeUserRole(ctx echo.Context, id string, newRole models.Role) (*models.User, error)
	GetUserByIdAuth(ctx echo.Context, idAuth string) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

var byId = "id_auth = ?"

func (d *UserRepositoryImpl) GetAllUsers(ctx echo.Context) ([]models.User, error) {
	var users []models.User
	if err := d.DB.WithContext(ctx.Request().Context()).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserRepositoryImpl) CreateUser(ctx echo.Context, user *models.User) (*models.User, error) {
	if err := d.DB.WithContext(ctx.Request().Context()).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserRepositoryImpl) UpdateUser(ctx echo.Context, id string, updatedUser *models.User) (*models.User, error) {
	if err := d.DB.WithContext(ctx.Request().Context()).Model(&models.User{}).Where(byId, id).Updates(updatedUser).Error; err != nil {
		return nil, err
	}

	// Devolver el usuario actualizado
	return updatedUser, nil
}

func (d *UserRepositoryImpl) GetUserByIdAuth(ctx echo.Context, idAuth string) (*models.User, error) {
	var user *models.User

	err := d.DB.WithContext(ctx.Request().Context()).Where(byId, idAuth).First(&user).Error

	return user, err
}
func (d *UserRepositoryImpl) ChangeUserRole(ctx echo.Context, id string, newRole models.Role) (*models.User, error) {
	var user models.User
	result := d.DB.WithContext(ctx.Request().Context()).Model(&user).Where(byId, id).Update("role", newRole)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
