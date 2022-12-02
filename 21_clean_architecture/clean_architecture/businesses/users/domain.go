package users

import (
	"gorm.io/gorm"
)

type Domain struct {
	*gorm.Model
  ID uint `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Usecase interface {
	GetAll() []Domain
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) string
}
type Repository interface {
	GetAll() []Domain
	Register(userDomain *Domain) Domain
	GetByEmail(userDomain *Domain) Domain
}