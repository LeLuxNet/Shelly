package main

import "io"

type Cmd interface {
	Run(args []string, session *Session) CmdCrashError
}

type InOutErr struct {
	In   io.Reader
	Out  io.WriteCloser
	Err  io.Writer
	Echo bool
}

func NewInOutErr(in io.Reader, out io.WriteCloser, err io.Writer, echo bool) InOutErr {
	if err == nil {
		return InOutErr{In: in, Out: out, Err: out, Echo: echo}
	}
	return InOutErr{In: in, Out: out, Err: err, Echo: echo}
}
