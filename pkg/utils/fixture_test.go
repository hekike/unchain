package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type TestRunner struct{}

func (c TestRunner) Run(
	dir string,
	command string,
	args ...string,
) ([]byte, error) {
	cs := []string{"-test.run=TestRunnerHelper", "--"}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Dir = dir
	cmd.Env = []string{
		"BUMP_ARGS=" + strings.Join(args, ","),
		"BUMP_DIR=" + dir,
	}
	out, err := cmd.CombinedOutput()
	return out, err
}

func TestRunnerHelper(*testing.T) {
	if os.Getenv("BUMP_ARGS") == "" {
		return
	}
	if os.Getenv("BUMP_DIR") == "" {
		return
	}
	defer os.Exit(0)
	fmt.Println(os.Getenv("BUMP_DIR") + "," + os.Getenv("BUMP_ARGS"))
}
