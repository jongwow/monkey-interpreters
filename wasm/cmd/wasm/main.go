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
func EchoWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid args"
		}
		ss := args[0].String()
		return ss
	})
}

func main() {
	fmt.Println("Hello Web Assembly from Go\n")

	js.Global().Set("getHtml", EchoWrapper()) // JS 로 편입 전에;.
	<-make(chan bool)
}
