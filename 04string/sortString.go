package main

import (
	"fmt"
	"strings"
)

func main() {

	// sl := []string{"a", "A", "B", "B", "d", "Z"}

	sliceString := []string{"Afzal", "Praveen", "Mayank", "Shubham", "Arif"}
	sliceString1 := make([]string, len(sliceString))
	var j int
	for i := 0; i < len(sliceString); i++ {
		sliceString1[j] = strings.ToUpper(sliceString[i])
		j++
	}
	fmt.Println(sliceString1)

	// str := "Afzal"
	// fmt.Println(strings.ToUpper(str))

	// sort.Strings(sl)
	// fmt.Println(sl)

	// sort.Strings(sliceString)
	// fmt.Println(sliceString)

}
