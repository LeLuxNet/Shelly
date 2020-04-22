package main

type Cmd interface {
	Run(args []string, session *Session) CmdCrashError
}
