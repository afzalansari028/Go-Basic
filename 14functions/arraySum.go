package main

import "fmt"

func main() {

	s := []int{10, 20, 30}
	fmt.Println(sum(s))

}

func sum(x []int) int {
	total := 0
	for _, val := range x {
		total = total + val
	}
	return total
}
