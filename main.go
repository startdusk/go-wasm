package main

import (
	"fmt"

	"syscall/js"
)

func main() {
	// 1.
	// copy wasm_exec.js to the project
	// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"

	// 2.
	// build wasm
	// GOOS=js GOARCH=wasm go build -o main.wasm

	// 3.
	// run http server
	// using goexec
	// go get -u github.com/shurcool/goexec
	// goexec 'http.ListenAndServe(":9080", http.FileServer(http.Dir(".")))'
	fmt.Println("Hello WebAssembly!")
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	p.Set("className", "block")

	styles := document.Call("createElement", "style")
	styles.Set("innerHTML", `
		.block {
			border: 1px solid black; color: white; background: black;
		}	
	`)

	document.Get("head").Call("appendChild", styles)
	document.Get("body").Call("appendChild", p)
}
