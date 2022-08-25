package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter any decimal value")
	var val int64
	fmt.Scanln(&val)

	output := strconv.FormatInt(val, 2)
	fmt.Println(output)

}
