package command_inpl

import (
	"github.com/LeLuxNet/Shelly/pkg/command"
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
	path, cErr := session.WorkingDir.GetRelativePath(args[1], true)
	if cErr != nil {
		return cErr
	}
	err := path.ExpectDir(false)
	if err != nil {
		return err
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
	return output.ClearScreen(std.Out)
}

type Mkdir struct{}

func (Mkdir) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	path, err := session.WorkingDir.GetRelativePath(args[1], false)
	if err != nil {
		return err
	}
	return os.Mkdir(path.General, os.ModePerm)
}

type Touch struct{}

func (Touch) Run(args []string, std sessions.Std, session *sessions.Session) error {
	if len(args) != 2 {
		return command.WrongArgCountError{Min: 1, Max: 1}
	}
	path, err := session.WorkingDir.GetRelativePath(args[1], false)
	if err != nil {
		return err
	}
	_, err = os.Stat(path.General)
	if os.IsNotExist(err) {
		file, err := os.Create(path.General)
		if err != nil {
			return err
		}
		return file.Close()
	}
	return path.Times(time.Now().Local())
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
	command.Register("mkdir", Mkdir{})
	command.Register("touch", Touch{})
	command.Register("run", Run{})

	command.Register("set", Set{})
	command.Register("export", Set{})

	command.Register("anf", Anf{})
}
