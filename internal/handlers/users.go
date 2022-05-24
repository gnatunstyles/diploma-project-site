package handlers

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"fmt"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	db.First(&user, id)
	// db.Find()
	if user.ID == 0 {
		return c.JSON(fiber.Map{"status": 404, "message": "user not found."})
	}
	return c.JSON(fiber.Map{"username": user.Username, "project": user.Project})
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	users := new([]models.User)
	db.Find(&users)
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	db := database.DBConn
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	user := new(models.User)
	db.Find(&user, id)
	return c.JSON(user)
}

func PostUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(503).SendString("Error. Wrong type of incoming data.")
	}
	db.Create(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	db.First(&user, id)
	if user.ID == 0 {
		return c.Status(500).SendString("User not found.")
	}
	db.Delete(&user)
	return c.JSON(&user)
}

func GetCurrentUser(c *fiber.Ctx) error {
	db := database.DBConn
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	user := models.User{}
	db.Where("id = ?", claims.Id).First(&user)

	return c.JSON(fiber.Map{
		"message":  "authorized",
		"user":     user,
		"username": user.Username,
		"claims":   claims,
	})
}

func UploadFile(c *fiber.Ctx) error {
	// id := 1
	file, err := c.FormFile(".laz")
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "File not found", "data": nil})
	}

	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "File not saved", "data": nil})
	}

	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)

	newProject := models.Project{
		UserId: 1,
		Name:   file.Filename,
		Size:   file.Size,
		Link:   "http://localhost:1234/examples/lion.html",
	}

	database.DBConn.Create(&newProject)

	return c.JSON(fiber.Map{"status": 500, "message": "File uploaded successfully."})

}

// func convert(file *multipart.FileHeader) error {
// 	return nil
// }
