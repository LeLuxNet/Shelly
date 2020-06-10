package output

import (
	"io"
)

const (
	NEWLINE = "\r\n"
)

func Send(msg string, writer io.Writer) {
	SendRaw([]byte(msg), writer)
}

func SendNl(msg string, writer io.Writer) {
	SendRaw([]byte(msg+NEWLINE), writer)
}

func SendRaw(raw []byte, writer io.Writer) {
	_, _ = writer.Write(raw)
}

func ClearScreen(out io.Writer) error {
	Send("\u001b[2J", out)
	return nil
}
