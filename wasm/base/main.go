package main

import (
	"syscall/js"
)

func greet(this js.Value, args []js.Value) any {
	js.Global().Get("document").Call("getElementById", "greet").Set("innerHTML", "HELLO!!")
	return nil
}

func main() {
	c := make(chan bool)
	js.Global().Set("goGreet", js.FuncOf(greet))
	<-c
}
