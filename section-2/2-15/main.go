package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/takepro14/section-2/section-2/2-15/hello"
)

func main() {
	x := hello.Input("input data")
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")

}
