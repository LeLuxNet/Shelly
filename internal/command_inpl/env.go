package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
)

type Set struct{}

func (Set) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) == 2 {
		return os.Setenv(args[1], "")
	}
	if len(args) != 3 {
		return command.WrongArgCountError{Min: 2, Max: 3}
	}
	return os.Setenv(args[1], args[2])
}
