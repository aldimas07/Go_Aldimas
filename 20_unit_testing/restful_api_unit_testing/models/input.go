package models

type LoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
