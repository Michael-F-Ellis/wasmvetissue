// +build js,wasm

package main

import "syscall/js"

var SP = &State{}

func main() {
	js.Global().Set("Foo", 1)
	Update()
}

func SetA(n int) {
	SP.A = n
}
