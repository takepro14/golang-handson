package main

import "fmt"

// Mydata is a structure.
type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{"Taro", []int{10, 20, 30}}
	hanako := Mydata{Name: "Hanako", Data: []int{90, 80, 70}}
	fmt.Println(taro)
	fmt.Println(hanako)
}
