package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"-"`
	ProjectNumber  uint64 `json:"project_number"`
	UsedSpace      uint64 `json:"used_space" gorm:"default:0"`
	AvailableSpace uint64 `json:"available" gorm:"default:10737418240"`
}

type Project struct {
	gorm.Model
	UserId   uint64 `json:"user_id"`
	Name     string `json:"project_name"`
	Info     string `json:"info" gorm:"default:'Simple point cloud with unknown user data.'"`
	Size     uint64 `json:"size"`
	Link     string `json:"link"`
	FilePath string `json:"file_path"`
	Points   uint64 `json:"points"`
}
