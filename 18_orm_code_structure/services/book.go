package services

import (
	database "Go_Aldimas/18_orm_code_structure/databases"
	models "Go_Aldimas/18_orm_code_structure/models"
	"errors"
)

func GetAllBook() ([]models.Book, error) {
	var book []models.Book = []models.Book{}

	if err := database.DB.Find(&book).Error; err != nil {
		return []models.Book{}, err
	}
	return book, nil
}

func GetBookByID(id string) (models.Book, error) {
	var book models.Book = models.Book{}

	if err := database.DB.Find(&book, "id = ?", id).Error; err != nil {
		return models.Book{}, err
	}

	if book.ID == 0 {
		return models.Book{}, errors.New("book not found")
	}

	return book, nil
}

func CreateBook(input models.Book) (models.Book, error) { //returnnya data buku yg baru ditambahkan. menangani req response bukan query db
	var book models.Book = models.Book{
		Name:        input.Name,
		Genre:       input.Genre,
		Description: input.Description,
	}

	if err := database.DB.Save(&book).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func UpdateBook(input models.Book, id string) (models.Book, error) {
	book, err := GetBookByID(id)

	if err != nil {
		return models.Book{}, nil
	}

	book.Name = input.Name
	book.Genre = input.Genre
	book.Description = input.Description

	if err := database.DB.Save(&book).Error; err != nil {
		return models.Book{}, nil
	}
	return book, nil
}

func DeleteBook(id string) bool {
	book, err := GetBookByID(id)

	if err != nil {
		return false
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return false
	}
	return true
}
