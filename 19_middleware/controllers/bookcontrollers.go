package controllers

import (
	model "Go_Aldimas/19_middleware/models"
	"Go_Aldimas/19_middleware/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBookController(c echo.Context) error {

	book, err := services.GetAllBook()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all book",
		"users":   book,
	})
}

func GetBookControllerID(c echo.Context) error {
	// your solution here
	var bookId string = c.Param("id")
	book, err := services.GetUserByID(bookId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "Book not found",
			"user":    book,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success get book data",
		"user":    book,
	})
}

func CreateBookController(c echo.Context) error {
	bookInput := model.Book{}
	c.Bind(&bookInput)

	book, err := services.CreateBook(bookInput)

	type BaseResponse struct {
    Message    string      `json:"message"`
    StatusCode int         `json:"200"`
    Data       interface{} `json:"book"`
}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK,BaseResponse{
		Message:    "Success Create Book",
		StatusCode: http.StatusOK,
		Data:       book,
}) 
	// return c.JSON(http.StatusOK, map[string]any{
	// 	"message": "success create new book",
	// 	"book":    book,
	// })
}

func DeleteBookController(c echo.Context) error {
	// your solution here
	var bookId string = c.Param("id")

	isDeleted := services.DeleteBook(bookId)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, map[string]any{
			"messages": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Book successfully deleted",
	})
}

func UpdateBookController(c echo.Context) error {
	// your solution here
	var bookId string = c.Param("id")
	var bookInput model.Book = model.Book{}
	c.Bind(&bookInput) //binding / parsing
	book, err := services.UpdateBook(bookInput, bookId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"Message": "Book Updated",
		"Book":    book,
	})

}
