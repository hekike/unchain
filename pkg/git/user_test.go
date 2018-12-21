package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	runner = TestRunner{}

	user, err := GetUser("../")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, &User{
		Name:  "John Doe",
		Email: "john.doe@test.com",
	}, user)
}

//////////// Fixtures

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
		"BUMP_ARGS=" + strings.Join(args, " "),
		"BUMP_DIR=" + dir,
	}
	out, err := cmd.CombinedOutput()
	return out, err
}

func TestRunnerHelper(*testing.T) {
	args := os.Getenv("BUMP_ARGS")
	dir := os.Getenv("BUMP_DIR")

	if args == "" {
		return
	}
	if dir == "" {
		return
	}

	if args == "config user.name" {
		fmt.Println("John Doe")
	} else if args == "config user.email" {
		fmt.Println("john.doe@test.com")
	} else {
		fmt.Println("no match: " + args)
	}
	defer os.Exit(0)
}
