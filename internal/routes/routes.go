package routes

import (
	"diploma-project-site/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

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
	app.Post("api/projects/update/:project_name", handlers.UpdateProject)
	app.Get("api/projects/share/:project_name", handlers.ShareProjectLink)
	app.Delete("api/projects/delete/:project_name", handlers.DeleteProject)

	app.Post("api/processing/random", handlers.RandomProcessingHandler)
	app.Post("api/processing/barycenter", handlers.GridBarycenterProcessingHandler)
	app.Post("api/processing/candidate", handlers.GridCandidateProcessingHandler)

}
