package handlers

import (
	database "diploma-project-site/db"
	b64 "encoding/base64"

	"strings"

	"diploma-project-site/internal/models"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *fiber.Ctx) error {
	db := database.DBConn
	header := c.Request().Header.Peek("authorization")
	user := new(models.User)
	str := strings.Split(string(header), " ")[1]
	uDec, err := b64.StdEncoding.DecodeString(str)
	if err != nil {
		return c.Status(400).SendString("Encoding error.")
	}

	creds := strings.Split(string(uDec), ":")
	req := &models.SignInRequest{
		Email:    creds[0],
		Password: creds[1],
	}

	db.Where("email = ?", req.Email).First(&user)
	if user.Email != "" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Wrong password. Try again!",
			})
		}

		token, err := generateJwt(c, *user)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error. Could not login to the server.",
			})
		}
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 6),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)
		return c.JSON(fiber.Map{
			"message": "success",
			"user":    user,
			"cookie":  cookie,
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Email not found. Try again!",
	})
}

func SignUp(c *fiber.Ctx) error {
	db := database.DBConn
	exist := new(models.User)
	header := c.Request().Header.Peek("authorization")
	str := strings.Split(string(header), " ")[1]
	uDec, err := b64.StdEncoding.DecodeString(str)

	if err != nil {
		return c.Status(400).SendString("Encoding error.")
	}

	creds := strings.Split(string(uDec), ":")
	req := &models.SignUpRequest{
		Email:    creds[0],
		Password: creds[1],
		Username: creds[2],
	}

	if req.Email == "" || req.Password == "" || req.Username == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Error. Sign-up credentials are wrong.")
	}

	db.Where("email = ?", req.Email).First(&exist)
	if exist.Email != "" {
		return fiber.NewError(fiber.StatusBadRequest, "Error. User with this email already exists.")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{Username: req.Username, Password: string(hash), Email: req.Email}
	db.Create(&user)

	return c.JSON(fiber.Map{"message": "registration success", "user": user})
}

func UserSignout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Second),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "user logged out successful"})
}

func generateJwt(c *fiber.Ctx, user models.User) (string, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
	})

	token, err := claims.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error. Could not login to the server.",
		})
	}
	return token, nil
}
