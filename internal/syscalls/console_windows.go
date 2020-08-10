// +build windows

package syscalls

import (
	"os"
	"syscall"
	"unsafe"
)

func isError(err error) bool {
	return err != nil && err.Error() != "The operation completed successfully."
}

func SetConsoleStdDefault() (error, error) {
	return SetConsoleStd(0x0200, 0x0001|0x0002|0x0004)
}

func SetConsoleStd(in uintptr, out uintptr) (error, error) {
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32DLL.NewProc("SetConsoleMode")

	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	inErr := setConsoleMode(in, os.Stdin.Fd(), proc)
	outErr := setConsoleMode(out, os.Stdout.Fd(), proc)
	return inErr, outErr
}

func GetConsoleStd() (uintptr, uintptr, error, error) {
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32DLL.NewProc("GetConsoleMode")

	in, inErr := getConsoleMode(os.Stdin.Fd(), proc)
	out, outErr := getConsoleMode(os.Stdout.Fd(), proc)
	return in, out, inErr, outErr
}

func setConsoleMode(data uintptr, handle uintptr, proc *syscall.LazyProc) error {
	_, _, err := proc.Call(handle, data)
	if isError(err) {
		return err
	}
	return nil
}

func getConsoleMode(handle uintptr, proc *syscall.LazyProc) (uintptr, error) {
	var old uint64 = 0
	_, _, err := proc.Call(handle, uintptr(unsafe.Pointer(&old)))
	if isError(err) {
		return 0, err
	}
	return uintptr(old), nil
}
