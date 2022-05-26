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
	projects := &[]models.Project{}
	db.Find(&projects)
	return c.JSON(projects)
}

func UpdateProject(c *fiber.Ctx) error {
	db := database.DBConn
	projectName := c.Params("project_name")

	p := &models.Project{}
	upd := &models.ProjectUpdateRequest{}

	db.Where("name = ?", projectName).First(&p)
	if p.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Error. Project with this ID not exists.",
			"data":    nil})
	}

	err := c.BodyParser(&upd)
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
		"message": "Project info updated successfully.",
		"project": p,
	})

}

func DeleteProject(c *fiber.Ctx) error {
	db := database.DBConn
	project := &models.Project{}
	projectName := c.Params("project_name")
	db.Where("name = ?", projectName).First(&project)

	if project.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Project not found.",
		})
	}

	user := &models.User{}
	db.First(&user, project.UserId)

	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error. User with this ID not found",
			"data":    nil})
	}
	err := os.RemoveAll(fmt.Sprintf("%s/%d/%s/", models.ProjectSavePath, user.ID, projectName))
	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error. File cannot be deleted.",
			"data":    nil})
	}

	user.AvailableSpace += uint64(project.Size)
	user.UsedSpace -= uint64(project.Size)
	user.ProjectNumber--

	db.Delete(&project)
	db.Save(&user)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Project was deleted successfully."})
}

func UploadProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	projectName := c.Params("project_name")

	number, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error during parsing id.",
			"data":    nil})
	}

	user := &models.User{}

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

	err = c.SaveFile(file, fmt.Sprintf("%s/%s/%s/%s", models.ProjectSavePath, id, projectName, file.Filename))
	if err != nil {
		cmd := exec.Command("ls")
		stdout, _ := cmd.Output()

		fmt.Println(string(stdout))
		err = os.Mkdir(fmt.Sprintf("%s/%s/%s/", models.ProjectSavePath, id, projectName), os.ModePerm)
		if err != nil {
			err = os.Mkdir(fmt.Sprintf("%s/%s/", models.ProjectSavePath, id), os.ModePerm)
			if err != nil {
				return c.JSON(fiber.Map{
					"status":  500,
					"message": "Cannot create user directory",
					"data":    nil,
					"error":   err.Error()})
			}
			_ = os.Mkdir(fmt.Sprintf("%s/%s/%s/", models.ProjectSavePath, id, projectName), os.ModePerm)
		}
		err = c.SaveFile(file, fmt.Sprintf("%s/%s/%s/%s", models.ProjectSavePath, id, projectName, file.Filename))
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

	project := &models.Project{}

	project.UserId = number
	project.Name = projectName
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
	user.ProjectNumber++

	fmt.Println(&user)
	db.Save(&user)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "File uploaded and converted successfully.",
		"project": project})

}

func ShareProjectLink(c *fiber.Ctx) error {
	db := database.DBConn
	pName := c.Params("project_name")
	p := &models.Project{}

	db.Where("name = ?", pName).First(&p)

	if p.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Error. Project with this name not exists.",
			"data":    nil})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Project link has been returned.",
		"link":    p.Link,
		"data":    nil})
}

func GetAllProjectsByUserId(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	number, err := strconv.ParseUint(string(id), 10, 64)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Error during parsing ID.",
			"data":    nil})
	}

	user := &models.User{}
	projectsList := &[]models.Project{}

	db.First(&user, number)
	if user.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Error. User with this ID not found.",
			"data":    nil})
	}

	db.Where("user_id = ?", user.ID).Find(&projectsList)
	if len(*projectsList) == 0 {
		return c.JSON(fiber.Map{
			"message":  "this user has no active projects",
			"projects": projectsList,
		})
	}
	return c.JSON(fiber.Map{
		"message":  "success",
		"projects": projectsList,
	})

}

func convert(c *fiber.Ctx, id uint, projectName string, f *multipart.FileHeader) (string, error) {
	inputRoot := fmt.Sprintf("%s/%d/%s/%s", models.ProjectSavePath, id, projectName, f.Filename)
	fmt.Println(inputRoot)
	outputDir := fmt.Sprintf("%s/%d/%s", models.ProjectSavePath, id, projectName)

	cmd := exec.Command(models.ConverterBuildPath, models.InputFlag, inputRoot, models.ProjectNameFlag, projectName, models.OutputFlag, outputDir)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	fmt.Println(string(stdout))
	link := fmt.Sprintf("%s%s/projects/%d/%s/%s.html", models.Host, models.PotreePort, id, projectName, projectName)
	return link, nil
}
