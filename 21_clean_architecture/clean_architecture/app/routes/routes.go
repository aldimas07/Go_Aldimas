package routes

import (
  "Go_Aldimas/21_clean_architecture/clean_architecture/controllers/users"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
  LoggerMiddleware echo.MiddlewareFunc
  JWTMiddleware    middleware.JWTConfig // JWT
	UserController   users.UserController // User
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// Logger
	e.Use(cl.LoggerMiddleware)
	v1 := e.Group("/v1")
	auth := v1.Group("/auth")
	// Login
	auth.POST("/login", cl.UserController.Login) // /v1/auth/login
	// SignUp
	auth.POST("/register", cl.UserController.Register) // v1/auth/register
	usersAdmin := v1.Group("/admin/users",middleware.JWTWithConfig(cl.JWTMiddleware)) // /v1/admins/users
	usersAdmin.GET("", cl.UserController.GetAll)
  //logout
  // withAuth := v1.Group("/auth", middleware.JWTWithConfig(cl.JWTMIddleware))
	// withAuth.POST("/logout", cl.UserController.Logout)
}