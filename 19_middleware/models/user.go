package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        //menandakan tabel user ditambahkan Created Updated Deleted at
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
}
type Book struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Genre       string `json:"genre" form:"genre"`
	Description string `json:"description" form:"description"`
}
