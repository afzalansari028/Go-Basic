package main

import "fmt"

func main() {
	//fmt.Println("hello world")

	msg := "hello-world"
	fmt.Println(alter(msg))
}

func alter(str string) string {
	str[5] = " "
	return str
}
