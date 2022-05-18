package main

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/handlers"
	"diploma-project-site/internal/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initRoutes(app *fiber.App) {
	// private := app.Group("/private")
	// public := app.Group("/public")
	app.Get("api/users", handlers.GetUsers)
	app.Get("api/users/:id", handlers.GetUserById)
	app.Post("api/users", handlers.PostUser)
	app.Delete("api/users/:id", handlers.DeleteUser)

	app.Post("api/sign-up", handlers.SignUp)
	app.Post("api/sign-in", handlers.SignIn)
	app.Get("api/user", handlers.GetCurrentUser)

	app.Post("api/users/:id/upload", handlers.UploadFile)

	app.Post("api/user/logout", handlers.UserSignout)

	// app.Post("/users/:id/:fileName/update", handlers.UpdateFile)
	// app.Delete("/users/:id/delete", handlers.DeleteFile)

	// private.Get("/")
}

func initDB() {
	var err error

	dsn := os.Getenv("DB_CONFIG_STRING")

	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	fmt.Println("database connected successfully")
	database.DBConn.AutoMigrate(&models.User{})
	fmt.Println("database migrated")

}

func main() {
	initDB()
	app := fiber.New(fiber.Config{
		BodyLimit: -1,
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //with this frontend allow to take cookie and send it back
	}))
	initRoutes(app)
	app.Listen(":8000")
}
