package routes

import (
	"auth-microservice/db"
	"auth-microservice/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)

}

func UpdateUserHandler(c echo.Context) error {
	id := c.Param("id")

	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return err
	}

	var existingUser models.User
	if err := db.DB.First(&existingUser, id).Error; err != nil {
		return err
	}

	existingUser.Name = updatedUser.Name
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password
	existingUser.Image = updatedUser.Image
	existingUser.Ubication = updatedUser.Ubication

	if err := db.DB.Save(&existingUser).Error; err != nil {
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

	// Validar que el nuevo rol sea uno de los valores permitidos
	switch newRole {
	case models.RoleAdmin, models.RoleUser, models.RoleEnterprise:
		// El nuevo rol es válido
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Rol no válido"})
	}

	// Buscar el usuario en la base de datos
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return err
	}

	// Actualizar el rol del usuario
	user.Role = newRole
	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserHandler(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, "No encontrado")
	}
	return c.JSON(http.StatusOK, user)

}
