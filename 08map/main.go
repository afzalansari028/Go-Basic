package main

import "fmt"

func main() {
	fmt.Println("--Welcome to Map in Golang--")

	var prodPrice map[string]int
	fmt.Println(prodPrice)

	tempPrice := make(map[string]int)
	tempPrice["convertible widget"] = 150

	prodPrice = tempPrice
	prodPrice["widget"] = 100
	fmt.Println(prodPrice)

	//for key, val := range prodPrice {
	//fmt.Printf("Key --> %v and value --> %v\n", key, val)
	//}

	//declare and initialization of map by literal
	empPrice := map[string]int{
		"widget": 75, "Orange": 50,
	}
	empPrice["Apple"] = 120
	fmt.Println(empPrice)

	fmt.Println(len(empPrice))
	el, ok := empPrice["widget"]
	fmt.Println(el, ok)
	delete(empPrice, "Orange")
	fmt.Println(len(empPrice))

}
