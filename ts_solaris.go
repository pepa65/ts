// +build solaris

package ts

import (
	"syscall"
	"golang.org/x/sys/unix"
)

const TIOCGWINSZ = 21608

func GetSize() (ws Size, err error) {
	ws = Size{80, 25, 0, 0}
	wsz, err := unix.IoctlGetWinsize(syscall.Stdout, TIOCGWINSZ)
	if err == nil {
		ws = Size{wsz.Row, wsz.Col, wsz.Xpixel, wsz.Ypixel}
	}
	return ws, err
}
