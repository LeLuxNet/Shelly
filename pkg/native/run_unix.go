// +build !windows

package native

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os/exec"
)

func Exec(args []string, std sessions.Std, dir string) error {
	_, err := exec.LookPath(args[0])
	if err != nil {
		return NoCmd{}
	}
	return execProgram(args[0], args[1:], std, dir)
}
