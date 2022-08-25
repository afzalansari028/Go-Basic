package main

import "fmt"

func main() {

	sl := []int{10, 20, 30, 40}
	for i := 0; i < len(sl); i++ {
		fmt.Println("Index :", i, " element :", sl[i])
	}

	fmt.Println()

	for _, value := range sl {
		fmt.Println("Index : ", "Value :", value)
	}
}
