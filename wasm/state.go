// +build js,wasm

package main

type State struct {
	A int
}

func Update() {
	SetA(5)
}
