package main

import "fmt"

func main() {

	if temp := -10; temp < 0 {
		fmt.Println("Below freezing")
	} else if temp == 0 {
		fmt.Println("At freezing")
	} else {
		fmt.Println("above freezing")
	}

}
