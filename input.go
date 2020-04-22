package main

import (
	"io"
	"os/exec"
	"runtime"
)

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

func (io InOutErr) ClearScreen() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = io.Out
	return cmd.Run()
}

func Send(msg string, writer io.Writer) {
	SendRaw([]byte(msg), writer)
}

func SendNl(msg string, writer io.Writer) {
	SendRaw([]byte(msg+Newline), writer)
}

func SendRaw(raw []byte, writer io.Writer) {
	_, _ = writer.Write(raw)
}
