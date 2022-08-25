package main

import "fmt"

func main() {
	/*	t := 4
		for t != 0 {
			fmt.Print("hi ")
			t--
		}   */

	var i = 0
	for {
		if i == 2 {
			break
		}

		fmt.Println(i)
		i++
	}
}
