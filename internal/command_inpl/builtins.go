package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/errors"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/parser"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io"
	"os"
	"strings"
	"time"
)

type Echo struct{}

func (Echo) Run(args []string, std sessions.Std, session *sessions.Session) error {
	output.SendNl(strings.Join(args[1:], " "), std.Out)
	return nil
}

type Cat struct{}

func (Cat) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	path, cErr := session.WorkingDir.GetRelativePath(args[1])
	if cErr != nil {
		return cErr
	}
	file, err := os.Open(path.General)
	if err != nil {
		return err
	}
	_, err = io.Copy(std.Out, file)
	if err != nil {
		return err
	}
	return nil
}

type Cd struct{}

func (Cd) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	return session.WorkingDir.ChangeDir(args[1])
}

type Exit struct{}

func (Exit) Run(args []string, std sessions.Std, session *sessions.Session) error {
	return session.Close()
}

type Pwd struct{}

func (Pwd) Run(args []string, std sessions.Std, session *sessions.Session) error {
	output.SendNl(session.WorkingDir.Visible, std.Out)
	return nil
}

type Sleep struct{}

func (Sleep) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	duration, err := parser.ParseTime(args[1])
	if err != nil {
		return err
	}
	time.Sleep(duration)
	return nil
}

type Clear struct{}

func (Clear) Run(args []string, std sessions.Std, session *sessions.Session) error {
	err := output.ClearScreen(std.Out)
	if err != nil {
		return errors.GeneralError{Message: err.Error()}
	}
	return nil
}

func Register() {
	command.Register("echo", Echo{})
	command.Register("cat", Cat{})
	command.Register("cd", Cd{})
	command.Register("ls", Ls{})

	command.Register("exit", Exit{})
	command.Register("quit", Exit{})

	command.Register("pwd", Pwd{})
	command.Register("sleep", Sleep{})
	command.Register("clear", Clear{})
	command.Register("run", Run{})

	command.Register("set", Set{})
	command.Register("export", Set{})
}
