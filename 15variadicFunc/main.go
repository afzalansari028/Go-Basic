package main

import "fmt"

func main() {
	vals := []int{1, 2, 3, 4}
	add(vals...)
	vals = []int{1, 2, 3, 4, 5, 6}
	add(vals...)

	add(1, 2, 3, 4)
	add(1, 2, 3, 4, 5, 6)

}

func add(vals ...int) {
	total := 0
	for _, val := range vals {
		total += val
	}
	fmt.Println(vals, total)
}
