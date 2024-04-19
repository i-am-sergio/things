package routes

import (
	"auth-microservice/controllers"
	"auth-microservice/middleware"

	"github.com/labstack/echo/v4"
)

func UsersRoutes(e *echo.Echo, uc *controllers.UserController) {
	users := e.Group("/users")
	users.GET("/", uc.GetAllUsersHandler, middleware.JWTMiddleware)
	users.POST("/", uc.CreateUserHandler, middleware.JWTMiddleware)
	users.GET("/:id", uc.GetUserHandler)
	users.PUT("/:id", uc.UpdateUserHandler, middleware.JWTMiddleware)
	users.PUT("/role/:id", uc.ChangeRoleHandler, middleware.JWTMiddleware)
}
