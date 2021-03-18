package shell

import (
	"os"
	"os/exec"
)

// Executes a command in an interactive shell. This does not return any of the
// output from the shell and is intended to be used when colorized output is
// desired. Can also be used to drop into SSH sessions when user control is
// needed.
func Interactive(command string) (err error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}

	cmd := exec.Command(shell, "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	return
}
