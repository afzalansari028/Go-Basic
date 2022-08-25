package main

import (
	"fmt"
)

func main() {

	x := "hello"

	y := &x

	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(x)
	*y = "goodbye"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*y)
}
