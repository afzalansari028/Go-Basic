package main

import "fmt"

func main() {
	n := 154
	var ld int
	var rev int
	for n != 0 {
		ld = n % 10
		rev = rev*10 + ld
		n = n / 10
	}
	fmt.Println(rev)
}
