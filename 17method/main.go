package main

import "fmt"

type Box struct {
	D float64
	B float64
	W float64
}

// func (b *Box) volume() float64 {      //--method
// 	vol := b.D * b.B * b.W
// 	return vol
// }

func volume(a, b, c float64) float64 { //--function
	vol := a * b * c
	return vol
}
func main() {
	// c := Box{D: 2, B: 4, W: 5}
	// fmt.Println(c.volume())

	var box Box
	box.B = 20
	box.W = 10
	box.D = 5
	fmt.Println(volume(box.B, box.D, box.W))
}
