package shell

import "os/exec"

// Execute a shell command. This command does not output anything to the terminal
// and should be used to retrieve raw output from the shell for processing. If
// colorized output is desired please use shell.Interactive.
func Execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()

	return string(out), err
}
