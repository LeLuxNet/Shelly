package output

import (
	"io"
	"os/exec"
	"runtime"
)

func Send(msg string, writer io.Writer) {
	SendRaw([]byte(msg), writer)
}

func SendNl(msg string, writer io.Writer) {
	SendRaw([]byte(msg+"\n"), writer)
}

func SendRaw(raw []byte, writer io.Writer) {
	_, _ = writer.Write(raw)
}

func ClearScreen(out io.Writer) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = out
	return cmd.Run()
}
