// +build windows

package native

import (
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"io/ioutil"
	"os"
	"os/exec"
)

func Exec(args []string, std sessions.Std, dir string) error {
	file, err := ioutil.TempFile(os.TempDir(), "exec_native.*.bat")
	if err != nil {
		return fallbackExec(args, std, dir)
	}
	defer os.Remove(file.Name())
	_, err = file.WriteString("@echo off\n" + join(args) + "\nexit /b %ERRORLEVEL%")
	if err != nil {
		return fallbackExec(args, std, dir)
	}

	err = execProgram("cmd.exe", append([]string{"/c"}, file.Name()), std, dir)
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 9009 {
		return NoCmd{}
	} else {
		return err
	}
}

func fallbackExec(args []string, std sessions.Std, dir string) error {
	return execProgram("cmd.exe", append([]string{"/c"}, args...), std, dir)
}
