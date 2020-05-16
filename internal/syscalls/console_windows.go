// +build windows

package syscalls

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func isError(err error) bool {
	return err != nil && err.Error() != "The operation completed successfully."
}

func SetConsoleStdDefault() {
	SetConsoleStd(0x0200, 0x0001|0x0002|0x0004)
}

func SetConsoleStd(in uintptr, out uintptr) {
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32DLL.NewProc("SetConsoleMode")

	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	setConsoleMode(in, os.Stdin.Fd(), proc)
	setConsoleMode(out, os.Stdout.Fd(), proc)
}

func GetConsoleStd() (uintptr, uintptr) {
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32DLL.NewProc("GetConsoleMode")

	in := getConsoleMode(os.Stdin.Fd(), proc)
	out := getConsoleMode(os.Stdout.Fd(), proc)
	return in, out
}

func setConsoleMode(data uintptr, handle uintptr, proc *syscall.LazyProc) {
	_, _, err := proc.Call(handle, data)
	if isError(err) {
		fmt.Println("Unable to set Windows Console mode: " + err.Error())
		os.Exit(1)
	}
}

func getConsoleMode(handle uintptr, proc *syscall.LazyProc) uintptr {
	var old uint64 = 0
	_, _, err := proc.Call(handle, uintptr(unsafe.Pointer(&old)))
	if isError(err) {
		fmt.Println("Unable to get Windows Console mode: " + err.Error())
		os.Exit(1)
	}
	return uintptr(old)
}
