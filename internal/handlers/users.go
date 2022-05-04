package handlers

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

func SignIn(c *fiber.Ctx) error {
	db := database.DBConn
	req := new(models.SignInRequest)
	user := new(models.User)
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(503).SendString("Error. Wrong type of incoming data.")
	}
	db.Where("email = ?", req.Email).First(&user)
	if user.Email != "" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			return c.Status(403).SendString("Wrong password. Try again!")
		}
		token, exp, err := generateJwt(*user)
		if err != nil {
			return err
		}
		return c.Status(200).JSON(fiber.Map{"token": token, "exp": exp, "user": user})
	}
	return c.Status(403).SendString("Email not found. Try again!")
}

func SignUp(c *fiber.Ctx) error {
	db := database.DBConn
	req := new(models.SignUpRequest)
	exist := new(models.User)
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(503).SendString("Error. Wrong type of incoming data.")
	}

	if req.Email == "" || req.Password == "" || req.Username == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Error. Sign-up credentials are wrong.")
	}

	db.Where("email = ?", req.Email).First(&exist)
	if exist.Email != "" {
		return fiber.NewError(fiber.StatusBadRequest, "Error. User with this email already exists.")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) // хэш
	if err != nil {
		return err
	}

	user := &models.User{Username: req.Username, Password: string(hash), Email: req.Email}
	db.Create(&user)

	token, exp, err := generateJwt(*user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func generateJwt(user models.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile(".las")
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "File not found", "data": nil})
	}
	return c.SaveFile(file, fmt.Sprintf("./projects/%s/%s", file.Filename, "user"))

}
