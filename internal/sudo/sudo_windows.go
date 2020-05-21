// +build windows

package sudo

import (
	"os"
	"syscall"
	"unsafe"
)

func IsRoot() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func runElevated(exe string, dir string, args string) error {
	dll := syscall.NewLazyDLL("shell32.dll")
	handle := dll.NewProc("ShellExecuteW")

	verbPtr, _ := syscall.UTF16PtrFromString("runas")
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(dir)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	_, _, err := syscall.Syscall6(handle.Addr(), 6, uintptr(0),
		uintptr(unsafe.Pointer(verbPtr)), uintptr(unsafe.Pointer(exePtr)),
		uintptr(unsafe.Pointer(cwdPtr)), uintptr(unsafe.Pointer(argPtr)), uintptr(1))
	if err != 0 {
		return err
	}
	return nil
}
