package services

import (
	"auth-microservice/models"
	"auth-microservice/repository"
	"context"
)

type UserService interface {
	GetAllUsersService(ctx context.Context) ([]models.User, error)
	CreateUserService(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUserService(ctx context.Context, id string, updatedUser *models.User) (*models.User, error)
	ChangeUserRoleService(ctx context.Context, id string, newRole models.Role) (*models.User, error)
	GetUserByIdAuthService(ctx context.Context, idAuth string) (*models.User, error)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetAllUsersService(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}
func (s *UserServiceImpl) CreateUserService(ctx context.Context, user *models.User) (*models.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserServiceImpl) UpdateUserService(ctx context.Context, id string, updatedUser *models.User) (*models.User, error) {
	return s.repo.UpdateUser(ctx, id, updatedUser)
}

func (s *UserServiceImpl) ChangeUserRoleService(ctx context.Context, idAuth string, newRole models.Role) (*models.User, error) {
	return s.repo.ChangeUserRole(ctx, idAuth, newRole)
}

func (s *UserServiceImpl) GetUserByIdAuthService(ctx context.Context, id string) (*models.User, error) {
	return s.repo.GetUserByIdAuth(ctx, id)
}
