# ts (Terminal Size)
Simple go application to get Terminal Size.
So many implementations do not support windows but `ts` has full windows support
(and OSX, Solaris, UNIX, etc.)

## Install cli
`go install github.com/pepa65/ts/cmd/ts@latest`

Or, after cloning the repo: `go install cmd/ts/ts.go`

### Usage
`ts`

Sample output: `80,25` (WIDTH,HEIGHT)

## Example use module
```go
package main

import (
	"fmt"
	"github.com/pepa65/ts"
)

func main() {
	term, _ := ts.GetSize()
	fmt.Printf("ColxRow: %dx%d  X,Y: %d,%d\n", term.W, term.H, term.X, term.Y)
}
```

### Issues
Getting the current position (.X and .Y) does not seem to work, at least on Linux.

## Documentation
[Original Documentation](http://godoc.org/github.com/olekukonko/ts)

* Changes: instead of calling `.Col()`, `.Row()`, `.PosX()` and `.PosY()` just access the members of the Size struct: `.W`, `.H`, `.X` and `.Y`
