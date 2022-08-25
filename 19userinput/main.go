package main

import "fmt"

func main() {
	fmt.Println("Enter your first name")
	var firstName string
	fmt.Scanln(&firstName)

	fmt.Println("Enter your second name")
	var secondName string
	fmt.Scanln(&secondName)

	fmt.Println("Your name is :", firstName, secondName)
}
