package ts

import (
	"fmt"
	"testing"
)

func ExampleGetSize() {
	size, _ := GetSize()
	fmt.Println(size.H)
	fmt.Println(size.W)
	fmt.Println(size.X)
	fmt.Println(size.Y)
}

func TestSize(t *testing.T) {
	size, err := GetSize()
	fmt.Println(size.H, size.W, size.X, size.Y)
	if err != nil {
		t.Fatal(err)
	}
	if size.W == 0 || size.H == 0 {
		t.Fatalf("Getting Terminal Size failed")
	}
}

