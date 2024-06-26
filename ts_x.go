// +build !windows,!solaris

package ts

import (
	"syscall"
	"unsafe"
)

func GetSize() (ws Size, err error) {
	_, _, ec := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)),
	)
	err = getError(ec)
	if TIOCGWINSZ == 0 && err != nil {
		ws = Size{25, 80, 0, 0}
	}
	return ws, err
}

func getError(ec interface{}) (err error) {
	switch v := ec.(type) {
	case syscall.Errno: // Some implementation return syscall.Errno number
		if v != 0 {
			err = syscall.Errno(v)
		}
	case error: // Some implementation return error
		err = ec.(error)
	default:
		err = nil
	}
	return err
}
