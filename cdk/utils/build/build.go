package build

import (
	"os"
	"os/exec"
)

func Golang(buildPath string, golangPath string) error {
	simpleCmd := exec.Command("go", "build", "-tags", " lambda.norpc", "-o", buildPath, golangPath)
	simpleCmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64", "CGO_ENABLED=0")
	_, err := simpleCmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
