package main

import "fmt"

type Box struct {
	Depth  float64
	Width  float64
	Height int32
}

func main() {
	b := Box{6, 3, 5}
	fmt.Println(b)
	c := Box{Depth: 5, Width: 4, Height: 3}
	c.Depth = 7
	ptr := &c
	(*ptr).Depth = 9
	fmt.Println(c)

	fmt.Printf("%+v", c)
}
