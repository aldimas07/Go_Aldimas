package services

//business logic mengenai data user

import (
	database "Go_Aldimas/18_orm_code_structure/databases"
	models "Go_Aldimas/18_orm_code_structure/models"
	"errors"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User = []models.User{}

	if err := database.DB.Find(&users).Error; err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func GetUserByID(id string) (models.User, error) {
	var user models.User = models.User{}

	if err := database.DB.Find(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	if user.ID == 0 {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func CreateUser(input models.User) (models.User, error) { //returnnya data user yg baru ditambahkan. menangani req response bukan query db
	// user := models.User{}
	var user models.User = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(input models.User, id string) (models.User, error) {
	user, err := GetUserByID(id)

	if err != nil {
		return models.User{}, nil
	}

	user.Email = input.Email
	user.Name = input.Name
	user.Password = input.Password

	if err := database.DB.Save(&user).Error; err != nil {
		return models.User{}, nil
	}
	return user, nil
}

func DeleteUser(id string) bool {
	user, err := GetUserByID(id)

	if err != nil {
		return false
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return false
	}
	return true
}
