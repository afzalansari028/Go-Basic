package main

import "fmt"

func main() {
	x := 20
	y := 10

	fmt.Println(add(x, y))
}

func add(x, y int) (a, b int, c bool) {

	a = x + y
	b = x - y
	c = x > y

	return
}
