package main

import (
	"embed"
	"fmt"
)

//go:embed 69-single_file.txt
var fileString string

//go:embed 69-single_file.txt
var fileByte []byte

//go:embed 69-single_file.txt
//go:embed 69-*.hash
var folder embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileByte), string(fileByte[:3]))

	content1, _ := folder.ReadFile("69-file1.hash")
	fmt.Println(string(content1))

	content2, _ := folder.ReadFile("69-file2.hash")
	fmt.Println(string(content2))
}
