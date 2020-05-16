package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
	"os"
)

type Set struct{}

func (Set) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *sessions.Session) error {
	if len(args) != 3 {
		return command.WrongArgCountError{Min: 2, Max: 2}
	}
	return os.Setenv(args[1], args[2])
}
