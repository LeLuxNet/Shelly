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
	defer syscalls.SetConsoleStd(in, out)

	input.ReaderInput(sessions.NewSession(os.Stdin, os.Stdout, os.Stderr, true))
}
