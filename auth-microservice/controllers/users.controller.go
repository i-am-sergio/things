package controllers

import (
	"auth-microservice/models"
	"auth-microservice/services"
	"auth-microservice/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token de autorización faltante")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return echo.NewHTTPError(http.StatusBadRequest, "Formato de token de autorización inválido")
	}

	token := authParts[1]
	fmt.Println(token)
	sub := utils.GetIdTokenJWTAuth0(token)
	var user models.User
	c.Bind(&user)
	user.IdAuth = sub
	createdUser, err := services.CreateUser(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func UpdateUserHandler(c echo.Context) error {
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
	user, err := services.UpdateUser(id, &updateUser)
	if err != nil {
		return err
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, "USER NOT FOUND")
	}
	return c.JSON(http.StatusOK, user)
}

func ChangeRoleHandler(c echo.Context) error {
	id := c.Param("id")

	// Obtener el nuevo rol del cuerpo de la solicitud
	var newRole models.Role
	if err := c.Bind(&newRole); err != nil {
		return err
	}

	user, err := services.ChangeUserRole(id, newRole)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	user, err := services.GetUserByIdAuth(id)

	if user == nil {
		return c.JSON(http.StatusNotFound, "USER NOT FOUND")
	}

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
