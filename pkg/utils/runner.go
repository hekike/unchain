package utils

import "os/exec"

// Runner runner interface
type Runner interface {
	Run(string, string, ...string) ([]byte, error)
}

// CommandRunner runs a command
type CommandRunner struct{}

// Run runs a command
func (c CommandRunner) Run(
	dir string,
	command string,
	args ...string,
) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	return cmd.CombinedOutput()
}
