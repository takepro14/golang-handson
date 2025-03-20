package main

import (
	"fmt"
	"strconv"

	"github.com/takepro14/section-2/2-11/hello"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print("1から" + x + "の合計は、")
	} else {
		return
	}
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println(t, "です。")
}
