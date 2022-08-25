package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	//formerly ioutil.ReadFile()
	//func ReadFile(name string)([]byte,error)
	contents, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(contents))
	fmt.Println(contents)
	fmt.Println(string(contents))
}
