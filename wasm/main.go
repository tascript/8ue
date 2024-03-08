package main

import (
	"fmt"
)

//go:wasmimport env add
func add(a int32) int32

func main() {
	fmt.Println(add(10))
}
