package main

import "fmt"

func main() {
	fmt.Println("Welcome to the function mechanism")
	sl := []float64{1.5, 1.2, 2.5, 6}
	res := avg(sl)
	fmt.Println("Average is :", res)

}

func avg(sl []float64) float64 {
	total := 0.0

	for _, val := range sl {
		total += val
	}
	return total / float64(len(sl))
}
