package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

// create file
// func main() {
// 	f, err := os.Create("create.txt")
// 	defer f.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(f)
// }

//open and close file
// func closer(f *os.File) error {
// 	f.Close()
// 	fmt.Println(f.Name(), "successfully closed")
// 	return nil
// }
// func main() {
// 	f, err := os.Open("file.txt")
// 	// defer closer(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("file successfully opened:", f.Name())
// }

//Remove or delete a file
// func main() {
// 	err := os.Remove("del.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("file successfully removed")
// }

//copy a file
// func main() {
// 	src, err := os.Open("src.txt")
// 	defer src.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//the flag allows the program to create it, if the file does not exist
// 	dst, err := os.OpenFile("dst.txt", os.O_RDWR|os.O_CREATE, 0755)
// 	defer dst.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//func Copy(dst Writer, src Reader) (written int64, err error)
// 	w, err := io.Copy(dst, src)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(reflect.TypeOf(w))
// 	fmt.Println(w)
// }

//Rename (move) file name
// func main() {
// 	oldPath := "file.txt"
// 	newPath := "./new/new.txt"
// 	//func Rename(oldpath, newpath string)error
// 	err := os.Rename(oldPath, newPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

//truncate a file
//myFile.txt originally contains "test file contents"
// func main() {
// 	//func Truncate(name string, size int64) error
// 	err := os.Truncate("myFile.txt", 10)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

//File Info
//stat returns a FileInfo that describes a file
//type FileInfo interface{
//	Name() string       --base name of the file
//  Size() int64        --length in bytes for regular files; system-dependent for others
//  Mode() FileMode     --file modes bits
//  ModeTime() time.Time --modification time
//  IsDir() bool        --abbreviation for Mode().IsDir()
//Sys() interface{}     --underlying data souurce (can return nil)
//}
// func main() {
// 	//func Stat(name string) (FileInfo, error)
// 	f, err := os.Stat("myFile.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("file name :", f.Name())
// 	log.Println("file size :", f.Size())
// 	log.Println("file permissions :", f.Mode())
// }
