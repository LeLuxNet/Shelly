// +build !windows

package main

import (
	"os"
	"os/exec"
)

func localInput() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	ReaderInput(NewInOutErr(os.Stdin, os.Stdout, os.Stderr, false))
}
