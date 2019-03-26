package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := "hello12"

	fmt.Print(unsafe.Sizeof(a))
}
