package drivers

import (
	userDomain "Go_Aldimas/22_docker/simpleapp/businesses/users"
	userDB "Go_Aldimas/22_docker/simpleapp/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
  return userDB.NewMySQLRepository(conn)
}