package main

import (
	"flag"
	"fmt"
)

/*
$ go run 72-command-line-flags.go -word=opt -numb=7 -fork -svar=flag
$ go run 72-command-line-flags.go -word=opt
$ go run 72-command-line-flags.go -word=opt a1 a2 a3
$ go run 72-command-line-flags.go -word=opt a1 a2 a3 -numb=7
$ go run 72-command-line-flags.go -h
$ go run 72-command-line-flags.go -wat
*/

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")

	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
