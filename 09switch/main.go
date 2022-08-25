package main

import "fmt"

func main() {

	workday := 6

	switch workday {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	//another way of usin switch statement
	temp := -10
	switch {

	case temp < 0:
		fmt.Println("below freezing")
	case temp > 0:
		fmt.Println("above freezing")
	default:
		fmt.Println("Freezing point")
	}

}
