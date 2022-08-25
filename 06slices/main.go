package main

import (
	"fmt"
)

func main() {

	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	//[]T[low:high]
	sl1 := sl[2:4]
	sl2 := sl[:5]
	sl3 := sl[3:]

	fmt.Println("values in sl1 :", sl1)
	fmt.Println("values in sl1 :", sl2)
	fmt.Println("values in sl1 :", sl3)

	var arr []string

	arr = append(arr, "afzal")
	arr = append(arr, "ansari")
	fmt.Println(arr)
}
