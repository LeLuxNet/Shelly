package main

type Cmd interface {
	Run(args []string, session *Session) CmdCrashError
}

var registeredCmds map[string]Cmd

func Register(listen string, cmd Cmd) bool {
	if _, ok := registeredCmds[listen]; ok {
		return false
	}
	registeredCmds[listen] = cmd
	return true
}
