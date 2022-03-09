package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = `<h4>Hello, I'm an HTML snippet from Go!`

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}

func main() {
	fmt.Println("Hello Web Assembly from Go\n")

	js.Global().Set("getHtml", GetHtml()) // JS 로 편입 전에;.
}
