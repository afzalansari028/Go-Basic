package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	x := 0
	y := len(arr) - 1

	for x < y {
		var temp = arr[x]
		arr[x] = arr[y]
		arr[y] = temp
		x++
		y--
	}
	fmt.Println(arr)
}
