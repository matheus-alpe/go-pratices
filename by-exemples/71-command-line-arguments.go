package main

import (
	"fmt"
	"os"
)

/*
$ go run 71-command-line-arguments.go a b c d e
*/

func main() {
	argsWithProgName := os.Args
	fmt.Println("argsWithProgName", argsWithProgName)

	argsWithoutProgName := argsWithProgName[1:]
	fmt.Println("argsWithoutProgName", argsWithoutProgName, len(argsWithoutProgName))

	if len(argsWithoutProgName) == 0 {
		panic("missing args")
	}

	firstArg := argsWithoutProgName[0]

	fmt.Println("firstArg", firstArg)
}
