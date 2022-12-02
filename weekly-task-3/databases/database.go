package database

import (
	"Go_Aldimas/20_unit_testing/restful_api_unit_testing/models"
	"errors"
	"fmt"
	"log"

	// "log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

var (
	DB *gorm.DB
)

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}

func InitTestDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "restfulapitest",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		// log.Fatalf("error while connecting to db: %s", err)
		panic(err)
	}

	// log.Println("connected to db")
	// DB.AutoMigrate(&models.Book{}, models.LoginInput{}, models.User{})
	InitialMigration()
}

func SeedUser()models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)

	var user models.User= models.User {
		Name: "unit test",
		Email: "unittest@gmail.com",
		Password : string(password),
	}

	if err := DB.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser models.User
	DB.Last(&createdUser)
	createdUser.Password = "123"


	return createdUser
}

func SeedBook()models.Book {

	var book models.Book= models.Book {
		Name: "unit test",
		Genre: "unit test genre",
		Description : "Description test",
	}

	if err := DB.Create(&book).Error; err != nil {
		panic(err)
	}

	var createdBook models.Book
	DB.Last(&createdBook)


	return createdBook
}

func CleanSeeders() {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	bookResult := DB.Exec("DELETE FROM books")
	userResult := DB.Exec("DELETE FROM users")

	var isFailed bool = bookResult.Error != nil || userResult.Error != nil

	if isFailed {
		panic(errors.New("error while clearing seeders"))
	}
	log.Println("Seeders cleaned up successfully")
}