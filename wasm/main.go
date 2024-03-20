package main

import (
	"syscall/js"
)

func greet(this js.Value, args []js.Value) any {
	js.Global().Get("document").Call("getElementById", "greet").Call("innerHTML", "HELLO!!")
	return nil
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("goGreet", js.FuncOf(greet))
	<-c
}
