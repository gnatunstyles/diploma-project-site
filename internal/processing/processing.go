package processing

import (
	"diploma-project-site/internal/models"
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
)

func ConvProcRand(projectName string, id uint, f *multipart.FileHeader, factor int) (string, error) {
	newProjDir := projectName + "_conv_rand/"
	fileToConvRoot := fmt.Sprintf("%s/%d/%s/%s", models.ProjectSavePath, id, projectName, f.Filename)
	outputDir := fmt.Sprintf("%s/%d/%s", models.ProjectSavePath, id, newProjDir)

	fileName := newProjDir[:len(newProjDir)-1]
	os.Mkdir(outputDir, os.ModePerm)

	cmd := exec.Command("python3", models.ProcessingRandomPath, fileToConvRoot, outputDir, fileName, string(factor))
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	fmt.Println(string(stdout))
	return "", nil
}

func ConvProcVox(projectName string, id uint, f *multipart.FileHeader, factor int) (string, error) {
	newProjDir := projectName + "_conv_rand/"
	fileToConvRoot := fmt.Sprintf("%s/%d/%s/%s", models.ProjectSavePath, id, projectName, f.Filename)
	outputDir := fmt.Sprintf("%s/%d/%s", models.ProjectSavePath, id, newProjDir)
	fileName := newProjDir[:len(newProjDir)-1]
	os.Mkdir(outputDir, os.ModePerm)
	cmd := exec.Command("python3", models.ProcessingRandomPath, fileToConvRoot, outputDir, fileName, string(factor))
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(string(stdout))
	return "", nil

} //rework

func ConvertPotreeProc() //todo
