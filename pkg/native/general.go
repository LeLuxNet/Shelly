package native

import (
	"io"
	"os/exec"
)

func execProgram(cmd string, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	exe := exec.Command(cmd, args...)
	exe.Stdin = stdin
	exe.Stdout = stdout
	exe.Stderr = stderr
	return exe.Run()
}
