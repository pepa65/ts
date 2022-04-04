// +build solaris

package ts

import (
	"syscall"
	"golang.org/x/sys/unix"
)

const (
	TIOCGWINSZ = 21608
)

// Get Windows Size
func GetSize() (ws Size, err error) {
	var wsz *unix.Winsize
	wsz, err = unix.IoctlGetWinsize(syscall.Stdout, TIOCGWINSZ)
	if err != nil {
		ws = Size{80, 25, 0, 0}
	} else {
		ws = Size{wsz.Row, wsz.Col, wsz.Xpixel, wsz.Ypixel}
	}
	return ws, err
}
