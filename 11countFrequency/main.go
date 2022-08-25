package main

import "fmt"

func main() {

	var str string = "abababaaabababdscccsscahck"
	var visited = make([]bool, len(str))

	for i := 0; i < len(str); i++ {
		count := 0

		if visited[i] == true {
			continue
		}

		for j := 0; j < len(str); j++ {
			if str[i] == str[j] {
				visited[j] = true
				count++
			}
		}
		fmt.Println(string(str[i]), ":", count)
	}
}
