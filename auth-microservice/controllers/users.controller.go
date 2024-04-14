package controllers

import (
	"auth-microservice/models"
	"auth-microservice/services"
	"auth-microservice/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		service: userService,
	}
}

func (uc *UserController) GetAllUsersHandler(c echo.Context) error {
	users, err := uc.service.GetAllUsersService(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUserHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")

	authParts := strings.Split(authHeader, " ")
	token := authParts[1]
	sub := utils.GetIdTokenJWTAuth0(token)
	var user models.User
	c.Bind(&user)
	user.IdAuth = sub
	createdUser, err := uc.service.CreateUserService(c, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) UpdateUserHandler(c echo.Context) error {
	id := c.Param("id")
	form, _ := c.MultipartForm()

	var updateUser models.User

	if name, ok := form.Value["name"]; ok && len(name) > 0 {
		updateUser.Name = name[0]
	}

	if email, ok := form.Value["email"]; ok && len(email) > 0 {
		updateUser.Email = email[0]
	}

	if password, ok := form.Value["password"]; ok && len(password) > 0 {
		updateUser.Password = password[0]
	}

	if ubication, ok := form.Value["ubication"]; ok && len(ubication) > 0 {
		updateUser.Ubication = ubication[0]
	}

	if file, err := c.FormFile("image"); err == nil {
		cloudinaryURL, err := utils.UploadImage(file)
		if err != nil {
			return err
		}
		updateUser.Image = cloudinaryURL
	}
	user, statusCode := uc.service.UpdateUserService(c, id, &updateUser)
	if statusCode != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) ChangeRoleHandler(c echo.Context) error {
	id := c.Param("id")

	var newRole models.Role
	c.Bind(&newRole)
	user, err := uc.service.ChangeUserRoleService(c, id, newRole)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	user, err := uc.service.GetUserByIdAuthService(c, id)

	if err != nil {
		return c.JSON(http.StatusNotFound, "USER NOT FOUND")
	}

	return c.JSON(http.StatusOK, user)
}
