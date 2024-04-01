package main

import (
	"fmt"

	"github.com/pepa65/ts/ts"
)

func main() {
	size, _ := ts.GetSize()
	fmt.Printf("%d,%d\n", size.W, size.H)
}
