package main // import "calc"

import (
	"strconv"
	"syscall/js"
)

var global = js.Global()

func add(this js.Value, i []js.Value) interface{} {
	value1 := i[0].String()
	value2 := i[1].String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	result := int1 + int2

	return result
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	global.Set("add", js.FuncOf(add))
	global.Set("add2", js.FuncOf(add))
	<-c
}
