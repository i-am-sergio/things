package services_test

import (
	"auth-microservice/models"
	"auth-microservice/services"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByIdAuth(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, id string, user *models.User) (*models.User, error) {
	args := m.Called(ctx, id, user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) ChangeUserRole(ctx context.Context, id string, newRole models.Role) (*models.User, error) {
	args := m.Called(ctx, id, newRole)
	return args.Get(0).(*models.User), args.Error(1)
}

func NewMockUserRepository() *MockUserRepository {
	return new(MockUserRepository)
}

func TestGelAllUsers(t *testing.T) {
	// GIVEN
	mockRepo := new(MockUserRepository)
	expectedUser := []models.User{
		{IdAuth: "1",
			Name:      "",
			Email:     "",
			Password:  "",
			Image:     "",
			Ubication: "",
			Role:      "ADMIN"},
	}

	service := services.NewUserService(mockRepo)
	mockRepo.On("GetAllUsers", mock.Anything).Return(expectedUser, nil)

	// WHEN
	result, err := service.GetAllUsersService(context.TODO())

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_Error(t *testing.T) {
	// GIVEN
	mockRepo := new(MockUserRepository)
	expectedError := errors.New("error fetching users")
	service := services.NewUserService(mockRepo)
	mockRepo.On("GetAllUsers", mock.Anything, mock.Anything).Return([]models.User{}, expectedError)

	// WHEN
	_, err := service.GetAllUsersService(context.TODO())

	// THEN
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByIdAuth(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      "ADMIN",
	}

	service := services.NewUserService(mockRepo)
	mockRepo.On("GetUserByIdAuth", mock.Anything, expectedUser.IdAuth).Return(expectedUser, nil)

	// WHEN
	result, err := service.GetUserByIdAuthService(context.TODO(), expectedUser.IdAuth)

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)

}

func TestGetUserByIdAuthError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      "ADMIN",
	}

	service := services.NewUserService(mockRepo)
	mockRepo.On("GetUserByIdAuth", mock.Anything, expectedUser.IdAuth).Return(&models.User{}, errors.New("error fetching user"))

	// WHEN
	_, err := service.GetUserByIdAuthService(context.TODO(), expectedUser.IdAuth)

	// THEN
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      "ADMIN",
	}
	service := services.NewUserService(mockRepo)
	mockRepo.On("CreateUser", mock.Anything, expectedUser).Return(expectedUser, nil)

	result, err := service.CreateUserService(context.TODO(), expectedUser)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)

}

func TestCreateUserError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      "ADMIN",
	}
	service := services.NewUserService(mockRepo)
	mockRepo.On("CreateUser", mock.Anything, expectedUser).Return(&models.User{}, errors.New("error creating user"))

	_, err := service.CreateUserService(context.TODO(), expectedUser)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "Pepe",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      "ADMIN",
	}
	service := services.NewUserService(mockRepo)
	mockRepo.On("UpdateUser", mock.Anything, expectedUser.IdAuth, expectedUser).Return(expectedUser, nil)

	result, err := service.UpdateUserService(context.TODO(), expectedUser.IdAuth, expectedUser)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "Pepe",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      "ADMIN",
	}
	service := services.NewUserService(mockRepo)
	mockRepo.On("UpdateUser", mock.Anything, expectedUser.IdAuth, expectedUser).Return(&models.User{}, errors.New("error updating user"))

	_, err := service.UpdateUserService(context.TODO(), expectedUser.IdAuth, expectedUser)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestChangeUserRole(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "Pepe",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      models.RoleAdmin,
	}
	service := services.NewUserService(mockRepo)
	mockRepo.On("ChangeUserRole", mock.Anything, expectedUser.IdAuth, models.RoleAdmin).Return(expectedUser, nil)

	result, err := service.ChangeUserRoleService(context.TODO(), expectedUser.IdAuth, models.RoleAdmin)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Role, result.Role)
	mockRepo.AssertExpectations(t)
}

func TestChangeUserRoleError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &models.User{
		IdAuth:    "1",
		Name:      "Pepe",
		Email:     "",
		Password:  "",
		Image:     "",
		Ubication: "",
		Role:      models.RoleAdmin,
	}
	service := services.NewUserService(mockRepo)
	mockRepo.On("ChangeUserRole", mock.Anything, expectedUser.IdAuth, models.RoleAdmin).Return(&models.User{}, errors.New("error changing role of user"))

	_, err := service.ChangeUserRoleService(context.TODO(), expectedUser.IdAuth, models.RoleAdmin)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
