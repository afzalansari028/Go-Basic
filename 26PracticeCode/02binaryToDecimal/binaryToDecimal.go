package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter test case")
	var t int
	fmt.Scanln(&t)
	for t != 0 {
		fmt.Println("Enter any binary value")
		var val string
		fmt.Scanln(&val)

		output, err := strconv.ParseInt(val, 2, 64)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
		t--

	}

}
