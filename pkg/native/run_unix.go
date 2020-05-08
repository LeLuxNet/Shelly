// +build !windows

package native

import (
	"io"
	"os/exec"
)

func Exec(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	_, err := exec.LookPath(args[0])
	if err != nil {
		return NoCmd{}
	}
	return execProgram(args[0], args[1:], stdin, stdout, stderr)
}
