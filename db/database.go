package database

import (
	"diploma-project-site/internal/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
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

func PlaceProcProjectToDB(id int, points uint64, fileName, newFilePath, link, prevProj, procType string) error {
	db := DBConn

	user := &models.User{}
	db.First(&user, id)
	if user.ID == 0 {
		return &fiber.Error{
			Code:    404,
			Message: "User not found."}
	}

	fileInfo, err := os.Stat(newFilePath)

	if err != nil {
		return err
	}

	project := &models.Project{
		UserId: uint64(id),
		Name:   fileName,
		Size:   uint64(fileInfo.Size()),
		Info: fmt.Sprintf("This point cloud was processed using %s algorithm. \nPrevious state of this cloud is %s project.",
			procType, prevProj),
		Link:     link,
		FilePath: newFilePath,
		Points:   uint64(points),
	}

	user.AvailableSpace -= project.Size
	user.UsedSpace += project.Size
	user.ProjectNumber++

	db.Save(&user)

	db.Create(&project)
	return nil
}
