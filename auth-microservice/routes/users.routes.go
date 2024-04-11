package routes

import (
	"auth-microservice/controllers"

	"github.com/labstack/echo/v4"
)

func UsersRoutes(e *echo.Echo, uc *controllers.UserController) {
	users := e.Group("/users")
	users.GET("/", uc.GetAllUsersHandler)
	// users.POST("/", uc.CreateUserHandler)
	// users.GET("/:id", uc.GetUserHandler, middleware.JWTMiddleware)
	// users.PUT("/:id", uc.UpdateUserHandler, middleware.JWTMiddleware)
	// users.PUT("/role/:id", uc.ChangeRoleHandler, middleware.JWTMiddleware)
}
