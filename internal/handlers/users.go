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

//ADD TO THE ENVIROMENTAL VARIABLES
const jwtSecretKey = "secret"

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
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Wrong password. Try again!",
			})
		}
		// token, exp, err := generateJwt(*user)
		// if err != nil {
		// 	return err
		// }

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Id:        strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
		})

		token, err := claims.SignedString([]byte(jwtSecretKey))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error. Could not login to the server.",
			})
		}
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true, //for storing into frontend and sending it
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

	return c.JSON(fiber.Map{"message": "registration success", "user": user})
}

func GetCurrentUser(c *fiber.Ctx) error {
	db := database.DBConn

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

func UserSignout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Second),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "user logged out successful"})
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile(".las")
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "File not found", "data": nil})
	}
	return c.SaveFile(file, fmt.Sprintf("./projects/%s/%s", file.Filename, "user"))

}

func GetProjects(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	db.First(&user, id)

	if user.ID == 0 {
		return c.JSON(fiber.Map{"status": 404, "message": "user not found."})
	}
	return c.JSON(fiber.Map{"username": user.Username, "projects": user.Projects})
}

func addProject(c *fiber.Ctx) error { return nil }

// func generateJwt(user models.User) (string, time.Time, error) {
// 	exp := time.Now().Add(time.Hour * 24) //1 day
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["user_id"] = user.ID
// 	claims["exp"] = exp
// 	t, err := token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		return "", time.Time{}, err
// 	}
// 	return t, exp, nil
// }
