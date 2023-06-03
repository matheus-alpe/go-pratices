package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)
	defer os.Remove(f.Name())
	fmt.Println("temp file name:", f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dirName, err := os.MkdirTemp("", "sample-dir")
	check(err)
	defer os.RemoveAll(dirName)
	fmt.Println("temp dir name:", dirName)

	filename := filepath.Join(dirName, "file1")
	err = os.WriteFile(filename, []byte{1, 2}, 0666)
	check(err)
	fmt.Println("temp filename path:", filename)
}
