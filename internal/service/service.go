package service

import (
	"diploma-project-site/internal/models"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

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
