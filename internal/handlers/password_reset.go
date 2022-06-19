package handlers

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func PasswordReset(c fiber.Ctx) error {
	db := database.DBConn
	req := &models.PasswordResetRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(
			fiber.Map{
				"status":  400,
				"message": "error. wrong request format",
				"error":   err,
			},
		)
	}
	tkn := generateUserToken(24) //models.ResetTokenLength
	_, _ = db, tkn
	return nil
}

func generateUserToken(length int) string {

	runes := []rune("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM") //models/getenv

	token := make([]rune, length)
	for i := range token {
		token[i] = runes[rand.Intn(len(runes))]
	}
	return string(token)
}
