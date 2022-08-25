package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	//func NewScanner(r io.Reader) *Scanner
	s := bufio.NewScanner(f)
	//func (s *Scan() bool
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
