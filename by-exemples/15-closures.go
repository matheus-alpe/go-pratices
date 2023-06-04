package main

import "fmt"

func sequencial() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	next := sequencial()
	fmt.Println(next())
	next()
	fmt.Println(next())

	next2 := sequencial()
	fmt.Println(next2())
}
