package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/engine"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io/ioutil"
)

type Run struct{}

func (Run) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	path, err := session.WorkingDir.GetRelativePath(args[1], true)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(path.General)
	if err != nil {
		return err
	}
	engine.MultiLineInput(string(data), std, session)
	return nil
}
