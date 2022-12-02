package response

import (
	"Go_Aldimas/21_clean_architecture/clean_architecture/businesses/users"
)

type User struct {
  ID uint `json:"id" gorm:"primaryKey"`
  Email string `json:"email"`
  Password string `json:"password"`
}

func FromDomain(domain users.Domain) User {
  return User{
    ID: domain.ID,
    Email: domain.Email,
    Password: domain.Password,
  }
}