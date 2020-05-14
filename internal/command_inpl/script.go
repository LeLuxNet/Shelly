package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/engine"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
	"io/ioutil"
)

type Run struct{}

func (Run) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	path, cErr := session.WorkingDir.GetRelativePath(args[1])
	if cErr != nil {
		return cErr
	}
	data, err := ioutil.ReadFile(path.General)
	if err != nil {
		return err
	}
	engine.MultiLineInput(string(data), session)
	return nil
}
