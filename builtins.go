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

func (o alwaysFail) Run(args []string, session *Session) CmdCrashError {
	return o.err
}

type Echo struct{}

func (Echo) Run(args []string, session *Session) CmdCrashError {
	if len(args) < 2 {
		return WrongArgCountError{min: 1}
	}
	io.WriteString(session.Out, strings.Join(args[1:], " ")+Newline)
	return nil
}

type Cat struct{}

func (Cat) Run(args []string, session *Session) CmdCrashError {
	if len(args) != 2 {
		return WrongArgCountError{min: 1, max: 1}
	}
	file, err := os.Open(args[1])
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
	return session.WorkingDir.ChDir(args[1])
}

type Ls struct{}

func (Ls) Run(args []string, session *Session) CmdCrashError {
	files, err := session.WorkingDir.ListDir(false)
	if err != nil {
		return GeneralError{Message: err.Error()}
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	session.Out.Write([]byte(strings.Join(result, " ") + Newline))
	return nil
}

type Exit struct{}

func (Exit) Run(args []string, session *Session) CmdCrashError {
	session.Close()
	return nil
}

type Pwd struct{}

func (Pwd) Run(args []string, session *Session) CmdCrashError {
	session.Out.Write([]byte(session.WorkingDir.Visible + Newline))
	return nil
}

type Sleep struct{}

func (Sleep) Run(args []string, session *Session) CmdCrashError {
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

func (Clear) Run(args []string, session *Session) CmdCrashError {
	session.ClearScreen()
	return nil
}
