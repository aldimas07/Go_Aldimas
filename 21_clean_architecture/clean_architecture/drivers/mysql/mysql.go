package mysql_driver

import (
	"Go_Aldimas/21_clean_architecture/clean_architecture/drivers/mysql/users"
	util "Go_Aldimas/21_clean_architecture/clean_architecture/utils"
	"fmt"
	"errors"
	"log"
	"gorm.io/driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}


func (config *ConfigDB) InitDB() *gorm.DB {
	// connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	config.DB_HOST,
	// 	config.DB_USERNAME,
	// 	config.DB_PASSWORD,
	// 	config.DB_NAME,
	// 	config.DB_PORT,
	// )
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	config.DB_USERNAME,
	config.DB_PASSWORD,
	config.DB_HOST,
	config.DB_PORT,
	config.DB_NAME,
)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to a database server: %s", err)
	}

	log.Println("connected to a database server")
	return db
}


func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.User{},
	)
}


func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("error when getting the database instance : %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection : %v", err)
		return err
	}
	log.Println("database connection is closed")
	return nil
}


func SeedUser(db *gorm.DB) users.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	fakeUser, _ := util.CreateFaker[users.User]()

	userRecord := users.User{
		Email:    fakeUser.Email,
		Password: string(password),
	}

	if err := db.Create(&userRecord).Error; err != nil {
		panic(err)
	}

	var foundUser users.User

	db.Last(&foundUser)

	foundUser.Password = "123123"

	return foundUser
}



func CleanSeeders(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	userResult := db.Exec("DELETE FROM users")

	var isFailed bool = userResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}
