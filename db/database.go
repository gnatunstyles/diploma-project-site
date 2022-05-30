package database

import (
	"diploma-project-site/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB(dsn string) {
	var err error

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	fmt.Println("database connected successfully")
	DBConn.AutoMigrate(&models.User{}, &models.Project{})
	fmt.Println("database migrated")
}
