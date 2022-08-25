package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Enter array size")
	var N int
	fmt.Scanln(&N)

	arr := make([]int, N)
	fmt.Println("Enter array elements")
	for i := 0; i < N; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Print("Original :", arr)
	fmt.Println()
	sort.Ints(arr)
	fmt.Print("Ascending sorted :", arr)
	fmt.Println()

	var res []int
	res = ReverseSort(arr)
	fmt.Print("Descending sorted :", res)
}

func ReverseSort(arr []int) []int {
	var sortedRev []int
	for i := len(arr) - 1; i >= 0; i-- {
		sortedRev = append(sortedRev, arr[i])
	}
	return sortedRev
}
