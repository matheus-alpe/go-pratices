package main

import (
	"fmt"
	"os"
	"strings"
)

/*
$ BAR=2 go run 74-env-vars.go
*/

func main() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
}
