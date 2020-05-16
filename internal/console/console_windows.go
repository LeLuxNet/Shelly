// +build windows

package console

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/internal/input"
	"github.com/LeLuxNet/Shelly/pkg/sessions"
	"os"
	"syscall"
	"unsafe"
)

func Local() {
	inHandle := syscall.Handle(os.Stdin.Fd())
	outHandle := syscall.Handle(os.Stdout.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	getConsoleModeProc := kernel32DLL.NewProc("GetConsoleMode")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")

	var oldIn uint64 = 0
	_, _, err := getConsoleModeProc.Call(uintptr(inHandle), uintptr(unsafe.Pointer(&oldIn)))
	if isError(err) {
		fmt.Println("Unable to get Windows Console mode: " + err.Error())
		os.Exit(1)
	}

	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	_, _, err = setConsoleModeProc.Call(uintptr(inHandle), 0x0200)
	if isError(err) {
		fmt.Println("Unable to set Windows Console mode: " + err.Error())
		os.Exit(1)
	}
	defer setConsoleModeProc.Call(uintptr(inHandle), uintptr(oldIn))

	var oldOut uint64 = 0
	_, _, err = getConsoleModeProc.Call(uintptr(outHandle), uintptr(unsafe.Pointer(&oldOut)))
	if isError(err) {
		fmt.Println("Unable to get Windows Console mode: " + err.Error())
		os.Exit(1)
	}

	_, _, err = setConsoleModeProc.Call(uintptr(outHandle), 0x0001|0x0002|0x0004)
	if isError(err) {
		fmt.Println("Unable to set Windows Console mode: " + err.Error())
		os.Exit(1)
	}
	defer setConsoleModeProc.Call(uintptr(outHandle), uintptr(oldOut))

	input.ReaderInput(sessions.NewSession(os.Stdin, os.Stdout, os.Stderr, true))
}

func isError(err error) bool {
	return err != nil && err.Error() != "The operation completed successfully."
}
