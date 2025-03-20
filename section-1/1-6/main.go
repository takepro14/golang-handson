package main

import (
	"fmt"

	"github.com/takepro14/section-1/1-6/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
