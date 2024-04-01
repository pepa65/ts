package main

import (
	"fmt"
	//"os"

	"github.com/pepa65/ts"
	//"golang.org/x/term"
)

func main() {
	size, _ := ts.GetSize()
	//fmt.Printf("WxH:%dx%d X:%d Y:%d\n", size.W, size.H, size.X, size.Y)
	fmt.Printf("%d,%d\n", size.W, size.H)
	//w, h, _ := term.GetSize(int(os.Stdout.Fd()))
	//fmt.Printf("%d,%d\n", w, h)
}
