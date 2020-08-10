// +build windows

package console

import (
	"github.com/LeLuxNet/Shelly/internal/input"
	"github.com/LeLuxNet/Shelly/internal/syscalls"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
)

func Local() {
	in, out, inErr, outErr := syscalls.GetConsoleStd()
	if inErr == nil && outErr == nil {
		defer syscalls.SetConsoleStd(in, out)
	}
	inErr, outErr = syscalls.SetConsoleStdDefault()
	echo := true
	if outErr != nil {
		echo = false
	}

	input.ReaderInput(sessions.NewSession(os.Stdin, os.Stdout, os.Stderr, echo, sessions.Local))
}
