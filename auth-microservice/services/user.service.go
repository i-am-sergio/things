package services

import (
	"auth-microservice/models"
	"auth-microservice/repository"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	GetAllUsersService(ctx echo.Context) ([]models.User, error)
	CreateUserService(ctx echo.Context, user *models.User) (*models.User, error)
	UpdateUserService(ctx echo.Context, id string, updatedUser *models.User) (*models.User, error)
	ChangeUserRoleService(ctx echo.Context, id string, newRole models.Role) (*models.User, error)
	GetUserByIdAuthService(ctx echo.Context, idAuth string) (*models.User, error)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetAllUsersService(ctx echo.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}
func (s *UserServiceImpl) CreateUserService(ctx echo.Context, user *models.User) (*models.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserServiceImpl) UpdateUserService(ctx echo.Context, id string, updatedUser *models.User) (*models.User, error) {
	return s.repo.UpdateUser(ctx, id, updatedUser)
}

func (s *UserServiceImpl) ChangeUserRoleService(ctx echo.Context, idAuth string, newRole models.Role) (*models.User, error) {
	return s.repo.ChangeUserRole(ctx, idAuth, newRole)
}

func (s *UserServiceImpl) GetUserByIdAuthService(ctx echo.Context, id string) (*models.User, error) {
	return s.repo.GetUserByIdAuth(ctx, id)
}
