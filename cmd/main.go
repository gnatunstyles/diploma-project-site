package main

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/handlers"
	"diploma-project-site/internal/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initRoutes(app *fiber.App) {
	private := app.Group("/private")
	public := app.Group("/public")
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUserById)
	app.Post("/users", handlers.PostUser)
	app.Delete("/users/:id", handlers.DeleteUser)
	public.Post("/sign-up", handlers.SignUp)
	public.Post("/sign-in", handlers.SignIn)
	private.Get("/")
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
	app := fiber.New()
	initDB()
	initRoutes(app)
	app.Listen(":3000")
}
