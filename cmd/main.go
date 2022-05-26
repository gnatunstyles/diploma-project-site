package main

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/config"
	"diploma-project-site/internal/handlers"
	"diploma-project-site/internal/models"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initRoutes(app *fiber.App) {

	app.Get("api/users", handlers.GetUsers)
	app.Get("api/users/:id", handlers.GetUserById)
	app.Post("api/users", handlers.PostUser)
	app.Delete("api/users/:id", handlers.DeleteUser)

	app.Post("api/sign-up", handlers.SignUp)
	app.Post("api/sign-in", handlers.SignIn)
	app.Get("api/user", handlers.GetCurrentUser)
	app.Post("api/user/logout", handlers.UserSignout)

	app.Get("api/projects", handlers.GetProjects)
	app.Get("api/projects/:id", handlers.GetAllProjectsByUserId)
	app.Post("api/projects/upload/:id/:project_name", handlers.UploadProject)
	app.Post("api/projects/update/:id", handlers.UpdateProject)
	app.Get("api/projects/share/:id", handlers.ShareProjectLink)
	app.Delete("api/projects/delete/:project_name", handlers.DeleteProject)

}

func initDB(dsn string) {
	var err error

	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	fmt.Println("database connected successfully")
	database.DBConn.AutoMigrate(&models.User{}, &models.Project{})
	fmt.Println("database migrated")
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Error during the config reading")
		return
	}

	initDB(cfg.DBConnString)
	app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024 * 1024,
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //with this frontend allow to take cookie and send it back
	}))

	initRoutes(app)
	app.Listen(models.BackendPort)
}
