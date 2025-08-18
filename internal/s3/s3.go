package s3

import (
	"fmt"
	"os"
	"os/exec"
)

const RJH = "myminio/rjh/"

func Run(name string, args ...string) {
	wd, _ := os.Getwd()

	cmd := exec.Command(name, args...)
	cmd.Dir = wd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Command failed: %v\n", err)
		os.Exit(1)
	}
}
