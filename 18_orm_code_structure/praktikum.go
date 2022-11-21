package main

import (
	"Go_Aldimas/18_orm_code_structure/controllers"
	database "Go_Aldimas/18_orm_code_structure/databases"

	"github.com/labstack/echo/v4"
)

func init() { //dieksekusi sebelum praktikum.go
	database.InitDB()           //konek ke db
	database.InitialMigration() //pembuatan tabel ke db
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetBookControllerID)
	e.POST("/books/", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("books/:id", controllers.UpdateBookController)

	

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":1323"))
}
