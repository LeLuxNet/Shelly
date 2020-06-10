// +build windows

package console

import (
	"github.com/LeLuxNet/Shelly/internal/input"
	"github.com/LeLuxNet/Shelly/internal/syscalls"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
)

func Local() {
	in, out := syscalls.GetConsoleStd()
	syscalls.SetConsoleStdDefault()

	input.ReaderInput(sessions.NewSession(os.Stdin, os.Stdout, os.Stderr, true))
	syscalls.SetConsoleStd(in, out)
}
