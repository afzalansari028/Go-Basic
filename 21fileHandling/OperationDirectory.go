package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// MkdirAll create directory
func main() {
	//func Stat(name string) (FileInfo, error)
	d, err := os.Stat("subdir")
	fmt.Println("error returned by os.Stat() is :", err)
	if err == nil {
		fmt.Println(d.Mode(), d.IsDir())
		log.Fatal("file/directory name already exists")
	}
	if errors.Is(err, os.ErrNotExist) {
		//func MkdirAll(path string,perm FileMode) error
		err := os.MkdirAll("subdir", 0777)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("subdir directory created")
	}
}

//MkdirAll Nested directories
// func main() {
// 	//func Join(elem ...string) string
// 	p := filepath.Join("..test", "subdir1", "subdir2")
// 	err := os.MkdirAll(p, 0777)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(p, "nested directory created")
// }

// ReadDir list files
// func main() {
// 	// formerly ioutil.ReadDir()
// 	//func ReadDir(name string) ([]DirEntry,error)
// 	ls, err := os.ReadDir("../22fileHandling")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, f := range ls {
// 		fmt.Println(f.Name(), f.IsDir())
// 	}
// }
