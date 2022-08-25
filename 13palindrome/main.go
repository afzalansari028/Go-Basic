package main

import "fmt"

func main() {

	str := "abba"
	ans := reverse(str)
	if str == ans {
		fmt.Println("Plaindrome")
	} else {
		fmt.Println("Not Palidrome")
	}
}

func reverse(str string) string {
	var s = ""
	for i := len(str) - 1; i >= 0; i-- {
		s = s + string(str[i])
	}
	return s
}
