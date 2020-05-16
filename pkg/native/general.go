package native

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os/exec"
)

func execProgram(cmd string, args []string, std sessions.Std, dir string) error {
	exe := exec.Command(cmd, args...)
	exe.Stdin = std.In
	exe.Stdout = std.Out
	exe.Stderr = std.Err
	exe.Dir = dir
	return exe.Run()
}
