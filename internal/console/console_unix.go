// +build !windows

package console

import (
	"github.com/LeLuxNet/Shelly/internal/input"
	"github.com/LeLuxNet/Shelly/pkg/models"
	"os"
	"os/exec"
)

func Local() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	input.ReaderInput(models.NewSession(os.Stdin, os.Stdout, os.Stderr, false))
}
