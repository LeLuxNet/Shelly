// +build !windows

package console

import (
	"github.com/LeLuxNet/Shelly/internal/input"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
	"os/exec"
)

func Local() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
    defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	input.ReaderInput(sessions.NewSession(os.Stdin, os.Stdout, os.Stderr, true))
}