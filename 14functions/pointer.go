package main

import "fmt"

func main() {
	x := 5
	y := 10
	res := adder(&x, &y)
	fmt.Println(res)
}

func adder(x, y *int) int {
	return (*x + *y)
}
