package drivers

import (
	userDomain "Go_Aldimas/21_clean_architecture/clean_architecture/businesses/users"
	userDB "Go_Aldimas/21_clean_architecture/clean_architecture/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
  return userDB.NewMySQLRepository(conn)
}