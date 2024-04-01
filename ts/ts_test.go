package ts

import (
	"fmt"
	"testing"
)

func ExampleGetSize() {
	size, _ := GetSize()
	fmt.Println(size.W)
	fmt.Println(size.H)
	fmt.Println(size.X)
	fmt.Println(size.Y)
}

func TestSize(t *testing.T) {
	size, err := GetSize()
	if err != nil {
		t.Fatal(err)
	}
	if size.W == 0 || size.H == 0 {
		t.Fatalf("Getting Terminal Size failed")
	}
}

