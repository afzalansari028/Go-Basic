package main

import "fmt"

func main() {

	var i, j int = 10, 20
	fmt.Println("Before swapping", i, j)

	i, j = j, i

	/*	i = i + j
		j = i - j
		i = i - j */

	fmt.Println("After swapped ", i, j)
	fmt.Printf("After swapped i = %d and j = %d ", i, j)
}
