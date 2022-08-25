package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	// f, err := os.Open("newFilee.txt")
	// defer f.Close()

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("File opened successfully", f)

	if _, err := os.Open("myFile.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("file does not exist")
		} else {
			log.Println(err)
		}
		return
	}
	fmt.Println("File is successfully opened")

}
