package main

/*
Use tags to run the program. E.g. Execute:
$ go run -tags "basics methodsinterfaces" ./main.go
to run init from basics and methods_interfaces package.

Using tags allow us to build different versions of the program:
$ go build -o <output-name> -tags "basics generics"
*/
import (
	_ "go-pratices/tour/01-basics"
	_ "go-pratices/tour/02-methods_interfaces"
	_ "go-pratices/tour/03-generics"
	_ "go-pratices/tour/04-concurrency"
)

func main() {}
