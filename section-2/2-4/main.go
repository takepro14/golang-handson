package main

import (
	"fmt"
	"strconv"

	"github.com/takepro14/section-2/2-4/hello"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	fmt.Println(x + "は、")
	if n%2 == 0 {
		fmt.Println("偶数です。")
	} else {
		fmt.Println("奇数です。")
	}
}
