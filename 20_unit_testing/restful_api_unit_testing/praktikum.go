package main

import (
	"Go_Aldimas/20_unit_testing/restful_api_unit_testing/constants"
	"Go_Aldimas/20_unit_testing/restful_api_unit_testing/controllers"
	database "Go_Aldimas/20_unit_testing/restful_api_unit_testing/databases"
	"Go_Aldimas/20_unit_testing/restful_api_unit_testing/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() { //dieksekusi sebelum praktikum.go
	database.InitDB()           //konek ke db
	database.InitialMigration() //pembuatan tabel ke db
}

func main() {
	// create a new echo instance
	e := echo.New()

	middlewares.LogMiddleware(e)

	privateRoutes := e.Group("") //mengotentikasi
	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(constants.JWT_SECRET_KEY),
	}))
	// Route / to handler function
	privateRoutes.GET("/users", controllers.GetUsersController)
	privateRoutes.GET("/users/:id", controllers.GetUserController)
	privateRoutes.DELETE("/users/:id", controllers.DeleteUserController)
	privateRoutes.PUT("/users/:id", controllers.UpdateUserController)

	//register login
	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.Login)

	


	//book route
	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetBookControllerID)
	privateRoutes.POST("/books/", controllers.CreateBookController)
	privateRoutes.DELETE("/books/:id", controllers.DeleteBookController)
	privateRoutes.PUT("books/:id", controllers.UpdateBookController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
