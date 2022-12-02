package request

import (
	"Go_Aldimas/21_clean_architecture/clean_architecture/businesses/users"

	"github.com/go-playground/validator"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserLogin) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}
