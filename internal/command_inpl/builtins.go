package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
	"github.com/LeLuxNet/Shelly/pkg/errors"
	"github.com/LeLuxNet/Shelly/pkg/output"
	"github.com/LeLuxNet/Shelly/pkg/parser"
	"github.com/LeLuxNet/Shelly/pkg/session"
	"io"
	"os"
	"strings"
	"time"
)

type Echo struct{}

func (Echo) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
	if len(args) < 2 {
		return command.WrongArgCountError{Min: 1}
	}
	output.SendNl(strings.Join(args[1:], " "), stdout)
	return nil
}

type Cat struct{}

func (Cat) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
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
	_, err = io.Copy(stdout, file)
	if err != nil {
		return err
	}
	return nil
}

type Cd struct{}

func (Cd) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	return session.WorkingDir.ChangeDir(args[1])
}

type Ls struct{}

func (Ls) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
	files, err := session.WorkingDir.ListDir(false)
	if err != nil {
		return err
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	output.SendNl(strings.Join(result, " "), stdout)
	return nil
}

type Exit struct{}

func (Exit) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
	return session.Close()
}

type Pwd struct{}

func (Pwd) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
	output.SendNl(session.WorkingDir.Visible, stdout)
	return nil
}

type Sleep struct{}

func (Sleep) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
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

func (Clear) Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, session *session.Session) error {
	err := output.ClearScreen(stdout)
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
}
