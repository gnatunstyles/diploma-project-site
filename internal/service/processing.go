package service

import (
	database "diploma-project-site/db"
	"diploma-project-site/internal/models"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ConvertProcRand(projectName, convFilePath string, id uint, factor int) (string, error) {
	processing := "random-sampling"
	newProjDir := projectName + "_rand/"
	fileToConvRoot := convFilePath
	outputDir := fmt.Sprintf("%s/%d/%s", models.ProjectSavePath, id, newProjDir)

	fileName := newProjDir[:len(newProjDir)-1]
	os.Mkdir(outputDir, os.ModePerm)
	fmt.Println("check")

	cmd := exec.Command(models.PythonVersion, models.ProcessingRandomPath, fileToConvRoot, outputDir, fileName, strconv.Itoa(factor))
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}
	fmt.Println(string(out))

	newFilePath := outputDir + fileName + ".las"

	link, err := ConvertNewProcPotree(int(id), newFilePath, fileName, outputDir)
	if err != nil {
		return "", err
	}

	points, err := GetPointsAmount(newFilePath, outputDir)
	if err != nil {
		return "", err
	}

	err = PlaceProcProjectToDB(int(id), points, fileName, newFilePath, link, projectName, processing)
	if err != nil {
		return "", err
	}

	return "Random sampling was done successfully.", nil
}

func ConvertProcCandidate(projectName, convFilePath string, id uint, voxelSize int) (string, error) {
	processing := "grid-center-candidate"

	newProjDir := projectName + "_cand/"
	fileToConvRoot := convFilePath
	outputDir := fmt.Sprintf("%s/%d/%s", models.ProjectSavePath, id, newProjDir)
	fileName := newProjDir[:len(newProjDir)-1]

	os.Mkdir(outputDir, os.ModePerm)
	fmt.Println("check")
	cmd := exec.Command(models.PythonVersion, models.ProcessingGridCandidatePath, fileToConvRoot, outputDir, fileName, strconv.Itoa(voxelSize))
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}
	fmt.Println(string(out))

	newFilePath := outputDir + fileName + ".las"

	link, err := ConvertNewProcPotree(int(id), newFilePath, fileName, outputDir)
	if err != nil {
		return "", err
	}

	points, err := GetPointsAmount(newFilePath, outputDir)
	if err != nil {
		return "", err
	}

	err = PlaceProcProjectToDB(int(id), points, fileName, newFilePath, link, projectName, processing)
	if err != nil {
		return "", err
	}

	return "Grid barycanter candidate algorithm sampling was done successfully.", nil

}

func ConvertProcBarycenter(projectName, convFilePath string, id uint, voxelSize int) (string, error) {
	processing := "grid-barycenter"
	newProjDir := projectName + "_bary/"
	fileToConvRoot := convFilePath
	outputDir := fmt.Sprintf("%s/%d/%s", models.ProjectSavePath, id, newProjDir)
	fileName := newProjDir[:len(newProjDir)-1]

	os.Mkdir(outputDir, os.ModePerm)
	fmt.Println("check")

	cmd := exec.Command(models.PythonVersion, models.ProcessingGridBarycenterPath, fileToConvRoot, outputDir, fileName, strconv.Itoa(voxelSize))
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}
	fmt.Println(string(out))

	newFilePath := outputDir + fileName + ".las"

	link, err := ConvertNewProcPotree(int(id), newFilePath, fileName, outputDir)
	if err != nil {
		return "", err
	}

	points, err := GetPointsAmount(newFilePath, outputDir)
	if err != nil {
		return "", err
	}

	err = PlaceProcProjectToDB(int(id), points, fileName, newFilePath, link, projectName, processing)
	if err != nil {
		return "", err
	}

	return "Grid barycenter algorithm sampling was done successfully.", nil

}

func ConvertNewProcPotree(id int, newFilePath, fileName, outputDir string) (string, error) {
	cmd := exec.Command(models.ConverterBuildPath, models.InputFlag, newFilePath, models.ProjectNameFlag, fileName, models.OutputFlag, outputDir)
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	fmt.Println(string(out))
	link := fmt.Sprintf("%s%s/%s/%d/%s/%s.html", models.PotreeHost, models.PotreePort, models.ProjectsDir, id, fileName, fileName)
	return link, nil
}

func PlaceProcProjectToDB(id int, points uint64, fileName, newFilePath, link, prevProj, procType string) error {
	db := database.DBConn

	user := &models.User{}
	db.First(&user, id)
	if user.ID == 0 {
		return &fiber.Error{
			Code:    404,
			Message: "User not found."}
	}

	fileInfo, err := os.Stat(newFilePath)

	if err != nil {
		return err
	}

	project := &models.Project{
		UserId: uint64(id),
		Name:   fileName,
		Size:   uint64(fileInfo.Size()),
		Info: fmt.Sprintf("This point cloud was processed using %s algorithm. \nPrevious state of this cloud is %s project.",
			procType, prevProj),
		Link:     link,
		FilePath: newFilePath,
		Points:   uint64(points),
	}

	user.AvailableSpace -= project.Size
	user.UsedSpace += project.Size
	user.ProjectNumber++

	db.Save(&user)

	db.Create(&project)
	return nil
}

func GetPointsAmount(filePath, outputDir string) (uint64, error) {
	cmd := exec.Command(models.PythonVersion, models.GetPointsAmountPath, filePath, outputDir, models.PointsFileName)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	fmt.Println(string(out))

	path := outputDir + models.PointsFileName

	buf, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return 0, err
	}

	pointsNum, err := strconv.Atoi(string(buf))
	if err != nil {
		return 0, err
	}

	err = os.Remove(path)
	if err != nil {
		return 0, err
	}

	return uint64(pointsNum), nil
}

func ConvertPotreeUploaded(id uint, projectName string, f *multipart.FileHeader) (string, error) {
	inputRoot := fmt.Sprintf("%s/%d/%s/%s", models.ProjectSavePath, id, projectName, f.Filename)
	fmt.Println(inputRoot)
	outputDir := fmt.Sprintf("%s/%d/%s", models.ProjectSavePath, id, projectName)

	cmd := exec.Command(models.ConverterBuildPath, models.InputFlag, inputRoot, models.ProjectNameFlag, projectName, models.OutputFlag, outputDir)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	fmt.Println(string(stdout))
	link := fmt.Sprintf("%s%s/projects/%d/%s/%s.html", models.PotreeHost, models.PotreePort, id, projectName, projectName)
	return link, nil
}
