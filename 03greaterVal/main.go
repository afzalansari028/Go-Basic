package main

import "fmt"

func main() {

	var a, b, c = 40, 30, 200

	if a > b && a > c {
		fmt.Println("Greater value is a =", a)
	} else if b > c {
		fmt.Println("Greater value is b =", b)
	} else {
		fmt.Println("Greater value is c =", c)
	}

}
