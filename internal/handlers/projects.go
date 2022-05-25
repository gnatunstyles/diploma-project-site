package handlers

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"fmt"
	"mime/multipart"
	"os"
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

func UpdateProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	p := new(models.Project)
	upd := new(models.ProjectUpdateRequest)

	number, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error during parsing ID.",
			"data":    nil})
	}

	user := new(models.User)

	db.First(&user, number)
	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Error. User with this ID not found.",
			"data":    nil})
	}

	db.First(&p, user.ProjectId)
	if p.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Error. Project with this ID not exists.",
			"data":    nil})
	}

	err = c.BodyParser(&upd)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error during parsing update request.",
			"data":    nil})
	}

	if upd.NewName != "" {
		p.Name = upd.NewName
	}

	if upd.NewInfo != "" {
		p.Info = upd.NewInfo
	}

	db.Save(&p)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Project info updated successfully."})

}

func DeleteProject(c *fiber.Ctx) error {
	db := database.DBConn
	project := new(models.Project)
	projectName := c.Params("project_name")
	db.Where("project_name = ?", projectName).First(&project)

	if project.Name == "" {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Project not found.",
		})
	}

	user := new(models.User)
	db.First(&user, project.UserId)

	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error. User with this ID not found",
			"data":    nil})
	}

	user.AvailableSpace += uint64(project.Size)
	user.UsedSpace -= uint64(project.Size)
	user.ProjectId = uint64(project.ID)

	db.Save(&user).Delete(&project)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Project was deleted successfully."})
}

func UploadProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	number, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error during parsing id.",
			"data":    nil})
	}

	user := new(models.User)

	db.First(&user, number)
	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error. User with this ID not found",
			"data":    nil})
	}

	file, err := c.FormFile("cloud")

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "File not found",
			"data":    nil})
	}

	if user.AvailableSpace < uint64(file.Size) {
		return c.JSON(fiber.Map{
			"status":  412,
			"message": "Error. Not enough space.",
			"data":    nil})
	}

	err = c.SaveFile(file, fmt.Sprintf("tools/potree/projects/%s/%s", id, file.Filename))
	if err != nil {
		err = os.Mkdir(fmt.Sprintf("tools/potree/projects/%s", id), os.ModePerm)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  500,
				"message": "Cannot create user directory",
				"data":    nil})
		}

		err = c.SaveFile(file, fmt.Sprintf("tools/potree/projects/%s/%s", id, file.Filename))
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  500,
				"message": "File not saved.",
				"data":    nil})
		}
	}

	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)

	project := new(models.Project)

	project.UserId = number
	project.Name = file.Filename
	project.Size = uint64(file.Size)

	project.Link, err = convert(c, uint(project.UserId), project.Name, file)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error during file converting: File not saved",
			"data":    nil,
			"error":   err})
	}

	fmt.Println(&project)

	fmt.Println("converted")

	db.Create(project)

	user.AvailableSpace -= uint64(project.Size)
	user.UsedSpace += uint64(project.Size)
	user.ProjectId = uint64(project.ID)
	fmt.Println(&user)
	db.Save(&user)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "File uploaded and converted successfully."})

}

func ShareProjectLink(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	p := new(models.Project)

	number, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error during parsing id.",
			"data":    nil})
	}

	user := new(models.User)

	db.First(&user, number)
	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Error. User with this ID not found.",
			"data":    nil})
	}

	db.First(&p, user.ProjectId)
	if p.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Error. Project with this ID not exists.",
			"data":    nil})
	}
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Project link has been returned.",
		"link":    p.Link,
		"data":    nil})
}

func convert(c *fiber.Ctx, id uint, projectName string, f *multipart.FileHeader) (string, error) {
	convRoute := "tools/PotreeConverter/build/PotreeConverter"
	inputRoot := fmt.Sprintf("tools/potree/projects/%d/%s", id, f.Filename) //todo for laz files
	fmt.Println(inputRoot)
	outputDir := fmt.Sprintf("tools/potree/projects/%d", id)

	cmd := exec.Command(convRoute, models.InputFlag, inputRoot, models.ProjectNameFlag, projectName, models.OutputFlag, outputDir)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	fmt.Println(string(stdout))

	link := fmt.Sprintf("%s%s/projects/%d/%s.html", models.Host, models.PotreePort, id, projectName)

	return link, nil
}
