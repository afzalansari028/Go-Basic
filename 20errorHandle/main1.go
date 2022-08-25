package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	radius := -1.0
	vol, err := volume(radius)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("voume of spere is : %0.3f", vol)
}

func volume(r float64) (float64, error) {
	if r < 0 {
		return 0, errors.New("Volume calculation failed; radius negative")
	}
	return (4.0 / 3.0) * 3.14 * r * r * r, nil
}

// func main() {
// 	val, err := CheckAge(20)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(val)
// }

// func CheckAge(age int) (string, error) {

// 	if age < 18 {
// 		return "", errors.New("Age should be more that 17")
// 	}
// 	return "You are eligible for vote", nil
// }
