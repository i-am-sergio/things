package routes

import (
	"auth-microservice/controllers"
	"auth-microservice/middleware"

	"github.com/labstack/echo/v4"
)

func UsersRoutes(e *echo.Echo) {
	users := e.Group("/users")

	users.POST("/", controllers.CreateUserHandler, middleware.JWTMiddleware)
	users.GET("/:id", controllers.GetUserHandler, middleware.JWTMiddleware)
	users.PUT("/:id", controllers.UpdateUserHandler, middleware.JWTMiddleware)
	users.PUT("/role/:id", controllers.ChangeRoleHandler, middleware.JWTMiddleware, middleware.RoleMiddleware)

}
