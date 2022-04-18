package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Project struct {
	Name string
	Link string
}
