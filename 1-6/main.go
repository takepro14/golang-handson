package main

import (
	"fmt"

	"github.com/takepro14/1-6/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
