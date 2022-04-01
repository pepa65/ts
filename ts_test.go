package ts

import (
	"fmt"
	"testing"
)

func ExampleGetSize() {
	size, _ := GetSize()
	fmt.Println(size.W) // Get Width
	fmt.Println(size.H) // Get Height
	fmt.Println(size.X) // Get X position
	fmt.Println(size.Y) // Get Y position
}

func TestSize(t *testing.T) {
	size, err := GetSize()

	if err != nil {
		t.Fatal(err)
	}
	if size.W == 0 || size.H == 0 {
		t.Fatalf("Screen Size Failed")
	}
}
