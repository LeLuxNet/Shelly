package main

import (
	"io"
	"os"
	"strings"
	"time"
)

type alwaysFail struct {
	err CmdCrashError
}

func (o alwaysFail) Run([]string, *Session) CmdCrashError {
	return o.err
}

type Echo struct{}

func (Echo) Run(args []string, session *Session) CmdCrashError {
	if len(args) < 2 {
		return WrongArgCountError{min: 1}
	}
	SendNl(strings.Join(args[1:], " "), session.Out)
	return nil
}

type Cat struct{}

func (Cat) Run(args []string, session *Session) CmdCrashError {
	if len(args) != 2 {
		return WrongArgCountError{min: 1, max: 1}
	}
	path, cErr := session.WorkingDir.GetRelativePath(args[1])
	if cErr != nil {
		return cErr
	}
	file, err := os.Open(path.General)
	if err != nil {
		return GeneralError{Message: err.Error()}
	}
	_, err = io.Copy(session.Out, file)
	if err != nil {
		return GeneralError{Message: err.Error()}
	}
	return nil
}

type Cd struct{}

func (Cd) Run(args []string, session *Session) CmdCrashError {
	if len(args) != 2 {
		return WrongArgCountError{min: 1, max: 1}
	}
	return session.WorkingDir.ChangeDir(args[1])
}

type Ls struct{}

func (Ls) Run(_ []string, session *Session) CmdCrashError {
	files, err := session.WorkingDir.ListDir(false)
	if err != nil {
		return GeneralError{Message: err.Error()}
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	SendNl(strings.Join(result, " "), session.Out)
	return nil
}

type Exit struct{}

func (Exit) Run(_ []string, session *Session) CmdCrashError {
	session.Close()
	return nil
}

type Pwd struct{}

func (Pwd) Run(_ []string, session *Session) CmdCrashError {
	SendNl(session.WorkingDir.Visible, session.Out)
	return nil
}

type Sleep struct{}

func (Sleep) Run(args []string, _ *Session) CmdCrashError {
	if len(args) != 2 {
		return WrongArgCountError{min: 1, max: 1}
	}
	duration, err := ParseTime(args[1])
	if err != nil {
		return err
	}
	time.Sleep(duration)
	return nil
}

type Clear struct{}

func (Clear) Run(_ []string, session *Session) CmdCrashError {
	err := session.ClearScreen()
	if err != nil {
		return GeneralError{Message: err.Error()}
	}
	return nil
}
