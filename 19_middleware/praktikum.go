package main

import (
	"Go_Aldimas/19_middleware/constants"
	"Go_Aldimas/19_middleware/controllers"
	database "Go_Aldimas/19_middleware/databases"
	"Go_Aldimas/19_middleware/middlewares"

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

	//register login
	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.Login)


	privateRoutes.DELETE("/users/:id", controllers.DeleteUserController)
	privateRoutes.PUT("/users/:id", controllers.UpdateUserController)

	
	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetBookControllerID)
	privateRoutes.POST("/books/", controllers.CreateBookController)
	privateRoutes.DELETE("/books/:id", controllers.DeleteBookController)
	privateRoutes.PUT("books/:id", controllers.UpdateBookController)

	

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":1323"))
}
