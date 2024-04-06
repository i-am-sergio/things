package controllers

import (
	"auth-microservice/models"
	"auth-microservice/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	createdUser, err := services.CreateUser(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func UpdateUserHandler(c echo.Context) error {
	id := c.Param("id")

	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return err
	}

	user, err := services.UpdateUser(id, &updatedUser)
	if user == nil {
		c.JSON(http.StatusNotFound, "USER NOT FOUND")
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "¡Usuario actualizado con éxito!")
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
	user, err := services.GetUser(id)

	if err != nil {
		return err
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, "USER NOT FOUND")
	}
	return c.JSON(http.StatusOK, user)
}
