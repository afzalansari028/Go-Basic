package main

import (
	"fmt"
)

func main() {

	var str = "golang"

	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
}
