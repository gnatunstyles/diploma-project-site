package handlers

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	db := database.DBConn
	projects := new([]models.Project)
	db.Find(&projects)
	return c.JSON(projects)
}

func UploadProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	number, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{"status": 50, "message": "Error during parsing id.", "data": nil})
	}

	file, err := c.FormFile(".las")
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "File not found", "data": nil})
	}

	err = c.SaveFile(file, fmt.Sprintf("tools/potree/projects/%s/%s", id, file.Filename))
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "File not saved", "data": nil})
	}

	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)

	newProject := models.Project{
		UserId: number,
		Name:   file.Filename,
		Size:   file.Size,
	}

	newProject.Link, err = convert(c, newProject.ID, newProject.Name)
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Error during file converting: File not saved", "data": nil, "error": err})
	}

	fmt.Println("converted")

	db.Create(&newProject)

	return c.JSON(fiber.Map{"status": 200, "message": "File uploaded and converted successfully."})

}

func DeleteProject(c *fiber.Ctx) error {
	db := database.DBConn
	project := new(models.Project)
	projectName := c.Params("project_name")
	db.Where("project_name = ?", projectName).First(&project)
	if project.Name == "" {
		return c.Status(500).SendString("Project not found.")
	}
	db.Delete(&project)
	return c.JSON(fiber.Map{"status": 200, "message": "Project was deleted successfully."})
}

func convert(c *fiber.Ctx, id uint, projectName string) (string, error) {
	convRoute := "tools/PotreeConverter/build/PotreeConverter"
	inputFlag := "-i"
	inputRoot := fmt.Sprintf("tools/potree/projects/%d/NEONDSSampleLiDARPointCloud.las", id)
	projectNameFlag := "-p"
	outputFlag := "-o"
	outputDir := fmt.Sprintf("tools/potree/projects/%d/%s", id, projectName)

	cmd := exec.Command(convRoute, inputFlag, inputRoot, projectNameFlag, projectName, outputFlag, outputDir)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	fmt.Println(string(stdout))

	link := fmt.Sprintf("%s/projects/%d/%s.html", hostName, id, projectName)

	return link, nil
}
