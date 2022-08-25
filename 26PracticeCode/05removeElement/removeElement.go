package main

import "fmt"

// func main() {
// 	// fmt.Println("Enter index value to remove")
// 	var index int
// 	// fmt.Scan(&index)

// 	sl := []int{1, 2, 3, 4, 5, 6}

// 	fmt.Printf("removed:%d\n", sl[index])
// 	sl = append(sl[:index], sl[index+1:]...)

// 	fmt.Println(sl)
// }

func main() {

	var sl1 []int
	sl := []int{1, 2, 3, 4, 5, 6}
	val := 2
	var removed int
	for i := 0; i < len(sl); i++ {
		if i == val {
			removed = sl[i]
			continue
		}
		sl1 = append(sl1, sl[i])
	}
	fmt.Println(sl)
	fmt.Println("removed :", removed)
	fmt.Println(sl1)
}
