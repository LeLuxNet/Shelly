// +build windows

package console

import (
	"fmt"
	"github.com/LeLuxNet/Shelly/internal/input"
	"github.com/LeLuxNet/Shelly/pkg/session"
	"os"
	"syscall"
	"unsafe"
)

func Local() {
	inHandle := syscall.Handle(os.Stdin.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	getConsoleModeProc := kernel32DLL.NewProc("GetConsoleMode")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")

	var old uint64 = 0
	_, _, err := getConsoleModeProc.Call(uintptr(inHandle), uintptr(unsafe.Pointer(&old)))
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

	defer setConsoleModeProc.Call(uintptr(inHandle), uintptr(old))

	outHandle := syscall.Handle(os.Stdout.Fd())
	_, _, err = setConsoleModeProc.Call(uintptr(outHandle), 0x0001|0x0002|0x0004)
	if isError(err) {
		fmt.Println("Unable to set Windows Console mode: " + err.Error())
		os.Exit(1)
	}

	input.ReaderInput(session.NewSession(os.Stdin, os.Stdout, os.Stderr, true))
}

func isError(err error) bool {
	return err != nil && err.Error() != "The operation completed successfully."
}
