package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string  `json:"email"`
	Username       string  `json:"username"`
	Password       string  `json:"-"`
	Project        Project `json:"project"`
	UsedSpace      uint64  `json:"used" gorm:"default:0"`
	AvailableSpace uint64  `json:"available" gorm:"default:10737418240"`
}

type Project struct {
	gorm.Model
	UserId uint64 `json:"user_id"`
	Name   string `json:"project_name"`
	Size   int64  `json:"size"`
	Link   string `json:"link"`
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
