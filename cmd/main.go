package main

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/config"
	"diploma-project-site/internal/models"
	"diploma-project-site/internal/routes"

	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal().Err(err).Msg("Error during the config reading")
		return
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //with this frontend allow to take cookie and send it back
	}))

	routes.InitRoutes(app)
	database.InitDB(cfg.DBConnString)

	// cer, err := tls.LoadX509KeyPair("")

	app.Listen(models.BackendPort)

}
