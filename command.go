package main

import (
	"io"
	"os/exec"
	"runtime"
)

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

func (io InOutErr) ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = io.Out
	cmd.Run()
}
