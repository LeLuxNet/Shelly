package main

import (
	"io/ioutil"
)

type Run struct{}

func (Run) Run(args []string, session *Session) CmdCrashError {
	if len(args) != 2 {
		return WrongArgCountError{min: 1, max: 1}
	}
	path, cErr := session.WorkingDir.GetRelativePath(args[1])
	if cErr != nil {
		return cErr
	}
	data, err := ioutil.ReadFile(path.General)
	if err != nil {
		return GeneralError{Message: err.Error()}
	}
	MultiLineInput(string(data), session)
	return nil
}
