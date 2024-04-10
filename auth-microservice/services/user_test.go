package services

// import (
// 	"auth-microservice/models"
// 	"net/http"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAllUsers(t *testing.T) {
// 	testCases := []struct {
// 		Name          string
// 		ExpectedError int
// 	}{
// 		{
// 			Name:          "List Users",
// 			ExpectedError: http.StatusOK,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			t.Parallel()
// 			con := &{}
// 			_, err := con.GetAllUsers()
// 			assert.Equal(t, tc.ExpectedError, err)
// 		})
// 	}
// }

// func TestCreateUser(t *testing.T) {
// 	testCases := []struct {
// 		Name          string
// 		User          *models.User
// 		ExpectedError int
// 	}{
// 		{
// 			Name:          "Create User",
// 			User:          &models.User{Name: "TestUser", IdAuth: "123"},
// 			ExpectedError: http.StatusOK,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			t.Parallel()
// 			con := &ServicesMocked{}
// 			_, err := con.CreateUser(tc.User)
// 			assert.Equal(t, tc.ExpectedError, err)
// 		})
// 	}
// }

// func TestUpdateUser(t *testing.T) {
// 	mock := &ServicesMocked{
// 		userMap: map[string]*models.User{
// 			"123": {IdAuth: "123", Name: "InitialUser"},
// 			// Agrega más usuarios iniciales si es necesario
// 		},
// 	}
// 	testCases := []struct {
// 		Name             string
// 		ID               string
// 		UpdatedUser      *models.User
// 		ExpectedError    int
// 		ExpectedUserData *models.User
// 	}{
// 		{
// 			Name:             "Update Name",
// 			ID:               "123",
// 			UpdatedUser:      &models.User{Name: "UpdatedName"},
// 			ExpectedError:    http.StatusOK,
// 			ExpectedUserData: &models.User{IdAuth: "123", Name: "UpdatedName"},
// 		},
// 		{
// 			Name:             "Update Email",
// 			ID:               "123",
// 			UpdatedUser:      &models.User{Email: "updated@example.com"},
// 			ExpectedError:    http.StatusOK,
// 			ExpectedUserData: &models.User{IdAuth: "123", Name: "UpdatedName", Email: "updated@example.com"},
// 		},
// 		// Agregar más casos de prueba según sea necesario
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			updatedUser, err := mock.UpdateUser(tc.ID, tc.UpdatedUser)

// 			// Verificar el resultado
// 			assert.Equal(t, tc.ExpectedError, err)
// 			if updatedUser != nil && tc.ExpectedUserData != nil {
// 				assert.Equal(t, *tc.ExpectedUserData, *updatedUser)
// 			}
// 		})
// 	}
// }

// func TestChangeUserRole(t *testing.T) {
// 	testCases := []struct {
// 		Name          string
// 		ID            string
// 		NewRole       models.Role
// 		ExpectedError int
// 	}{
// 		{
// 			Name:          "Change User Role",
// 			ID:            "123",
// 			NewRole:       models.RoleAdmin,
// 			ExpectedError: http.StatusOK,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			con := &ServicesMocked{}
// 			_, err := con.ChangeUserRole(tc.ID, tc.NewRole)
// 			assert.Equal(t, tc.ExpectedError, err)
// 		})
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	testCases := []struct {
// 		Name          string
// 		ID            string
// 		ExpectedError int
// 	}{
// 		{
// 			Name:          "Get User",
// 			ID:            "123",
// 			ExpectedError: http.StatusOK,
// 		},
// 		{
// 			Name:          "User that not exists",
// 			ID:            "1",
// 			ExpectedError: http.StatusNotFound,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.Name, func(t *testing.T) {
// 			t.Parallel()
// 			con := &ServicesMocked{}
// 			_, err := con.GetUser(tc.ID)
// 			assert.Equal(t, tc.ExpectedError, err)
// 		})
// 	}
// }
