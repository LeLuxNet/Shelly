package main

import (
	"io"
	"os"
	"strings"
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
	dir, err := os.Open(session.WorkingDir.General)
	if err != nil {
		return GeneralError{Message: err.Error()}
	}
	names, err := dir.Readdirnames(0)
	if err != nil {
		return GeneralError{Message: err.Error()}
	}
	session.Out.Write([]byte(strings.Join(names, " ") + Newline))
	return nil
}

type Exit struct{}

func (Exit) Run(args []string, session *Session) CmdCrashError {
	session.Close()
	return nil
}
