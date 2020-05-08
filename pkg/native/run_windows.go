// +build windows

package native

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Exec(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	file, err := ioutil.TempFile(os.TempDir(), "exec_native.*.bat")
	if err != nil {
		return fallbackExec(args, stdin, stdout, stderr)
	}
	defer os.Remove(file.Name())
	_, err = file.WriteString("@echo off\n" + strings.Join(args, " ") + "\nexit /b %ERRORLEVEL%")
	if err != nil {
		return fallbackExec(args, stdin, stdout, stderr)
	}

	err = execProgram("cmd.exe", append([]string{"/c"}, file.Name()), stdin, stdout, stderr)
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 9009 {
		return NoCmd{}
	} else {
		return err
	}
}

func fallbackExec(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	return execProgram("cmd.exe", append([]string{"/c"}, args...), stdin, stdout, stderr)
}
