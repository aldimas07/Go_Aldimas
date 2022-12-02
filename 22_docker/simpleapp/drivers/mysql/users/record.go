package users

import (
	"Go_Aldimas/22_docker/simpleapp/businesses/users"

	"gorm.io/gorm"
)

type User struct {
  *gorm.Model
  ID uint `json:"id" gorm:"size:100;primaryKey"`
  Email string `json:"email" gorm:"unique" faker:"email"`
  Password string `json:"password"`
}

func FromDomain(domain *users.Domain) *User {
  return &User {
    ID: domain.ID,
    Email: domain.Email,
    Password: domain.Password,
  }
}
func (rec *User) ToDomain() users.Domain{
  return users.Domain{
    ID: rec.ID,
    Email: rec.Email,
    Password: rec.Password,
  }
}