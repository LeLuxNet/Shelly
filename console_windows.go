// +build windows

package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
)

func localInput() {
	handle := windows.Handle(os.Stdin.Fd())

	var old uint32 = 0
	windows.GetConsoleMode(handle, &old)

	var mode uint32 = 0
	// https://docs.microsoft.com/en-us/windows/console/setconsolemode

	// mode |= windows.ENABLE_ECHO_INPUT // Needs ENABLE_LINE_INPUT
	// mode |= windows.ENABLE_PROCESSED_INPUT
	// mode |= windows.ENABLE_LINE_INPUT
	// mode |= windows.ENABLE_WINDOW_INPUT
	// mode |= windows.ENABLE_MOUSE_INPUT
	// mode |= windows.ENABLE_INSERT_MODE // Needs ENABLE_EXTENDED_FLAGS
	// mode |= windows.ENABLE_QUICK_EDIT_MODE // Needs ENABLE_EXTENDED_FLAGS
	// mode |= windows.ENABLE_EXTENDED_FLAGS
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_INPUT

	if err := windows.SetConsoleMode(handle, mode); err != nil {
		fmt.Println("Unable to set Windows Console mode: " + err.Error())
		os.Exit(1)
	}

	defer windows.SetConsoleMode(handle, old)

	ReaderInput(NewInOutErr(os.Stdin, os.Stdout, os.Stderr, true))
}
