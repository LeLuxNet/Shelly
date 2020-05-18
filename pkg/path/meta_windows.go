// +build windows

package path

import (
	"syscall"
	"time"
)

func (p Path) Times(time time.Time) error {
	path, err := syscall.UTF16PtrFromString(p.General)
	if err != nil {
		return err
	}
	handle, err := syscall.CreateFile(path,
		syscall.FILE_WRITE_ATTRIBUTES, syscall.FILE_SHARE_WRITE, nil,
		syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(handle)
	fileTime := syscall.NsecToFiletime(time.UnixNano())
	return syscall.SetFileTime(handle, &fileTime, &fileTime, &fileTime)
}
