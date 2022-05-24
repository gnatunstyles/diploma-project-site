package handlers

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const hostName string = "http://localhost:1234"

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
