package npm

import "os/exec"

// Runner runner interface
type Runner interface {
	Run(string, string, ...string) ([]byte, error)
}

type commandRunner struct{}

func (c commandRunner) Run(
	dir string,
	command string,
	args ...string,
) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	return cmd.CombinedOutput()
}
