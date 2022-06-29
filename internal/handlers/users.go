package handlers

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

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
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "error. wrong request format",
		})
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
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "error. wrong request format",
		})
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
		return c.JSON(fiber.Map{
			"status":  404,
			"message": "error. user not found.",
		})
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
		"message":    "authorized",
		"user":       user,
		"username":   user.Username,
		"email":      user.Email,
		"used_space": user.UsedSpace,
		"avaliable":  user.AvailableSpace,
		"claims":     claims,
	})
}

func EditCurrentUser(c *fiber.Ctx) error {
	db := database.DBConn
	upd := &models.UserUpdateRequest{}

	err := c.BodyParser(&upd)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  422,
			"message": "error. wrong type of incoming data",
		})
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	user := models.User{}
	db.Where("id = ?", claims.Id).First(&user)

	if upd.NewEmail != "" {
		dup := models.User{}
		db.Where("email = ?", upd.NewEmail).First(&dup)
		if dup.ID == 0 {
			user.Email = upd.NewEmail
		} else {
			return c.JSON(fiber.Map{
				"status":  fiber.StatusMethodNotAllowed,
				"message": "error. user with this email already exists.",
			})
		}
	}
	if upd.NewUsername != "" {
		{
			dup := models.User{}
			db.Where("username = ?", upd.NewUsername).First(&dup)
			if dup.ID == 0 {
				user.Username = upd.NewUsername
			} else {
				return c.JSON(fiber.Map{
					"status":  fiber.StatusMethodNotAllowed,
					"message": "error. user with this name already exists.",
				})
			}
		}
	}

	db.Save(&user)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "User info updated successfully.",
		"user":    user})

}
