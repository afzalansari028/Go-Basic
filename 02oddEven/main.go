package main

import "fmt"

func main() {

	num := 20

	if num%2 == 0 {
		fmt.Printf("%d is even number", num)
	} else {
		fmt.Printf("%d is odd number", num)
	}

}
