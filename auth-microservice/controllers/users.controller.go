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
	UserService services.UserServiceImpl
}

func NewUserController(userService services.UserServiceImpl) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) GetAllUsersHandler(c echo.Context) error {
	users, err := uc.UserService.GetAllUsersService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUserHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token de autorización faltante")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return echo.NewHTTPError(http.StatusBadRequest, "Formato de token de autorización inválido")
	}

	token := authParts[1]
	sub := utils.GetIdTokenJWTAuth0(token)
	var user models.User
	c.Bind(&user)
	user.IdAuth = sub
	createdUser, err := uc.UserService.CreateUserService(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) UpdateUserHandler(c echo.Context) error {
	id := c.Param("id")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
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
		cloudinaryURL, err := services.UploadImage(file)
		if err != nil {
			return err
		}
		updateUser.Image = cloudinaryURL
	}
	user, statusCode := uc.UserService.UpdateUserService(c.Request().Context(), id, &updateUser)
	if statusCode != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) ChangeRoleHandler(c echo.Context) error {
	id := c.Param("id")

	// Obtener el nuevo rol del cuerpo de la solicitud
	var newRole models.Role
	if err := c.Bind(&newRole); err != nil {
		return err
	}
	user, err := uc.UserService.ChangeUserRoleService(c.Request().Context(), id, newRole)
	if err == nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	user, err := uc.UserService.GetUserByIdAuthService(c.Request().Context(), id)

	if err == nil {
		return c.JSON(http.StatusNotFound, "USER NOT FOUND")
	}

	return c.JSON(http.StatusOK, user)
}
