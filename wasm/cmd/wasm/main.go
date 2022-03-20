package main

import (
	"fmt"
	"syscall/js"

	"github.com/jongwow/monkey/repl"
)

func InterpreterWrapper(val chan<- string) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid args"
		}
		ss := args[0].String()
		val <- ss
		return ss
	})
}

func main() {
	done := make(chan bool)
	inStr := make(chan string)
	outStr := make(chan string)
	fmt.Println("Hello Web Assembly from Go\n")

	go func() {
		for {
			select {
			case <-done:
				return
			case received := <-outStr:
				js.Global().Get("document").Call("getElementById", "interpretOutput").Set("innerText", received)
			}
		}
	}()
	go repl.StartByLine(done, inStr, outStr)
	js.Global().Set("monkeyInterpret", InterpreterWrapper(inStr)) // JS 로 편입 전에;.
	<-done
}
