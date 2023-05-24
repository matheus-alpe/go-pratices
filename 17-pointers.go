package main

import (
	"fmt"
)

func zeroThis(ival int) {
	ival = 0
}

func zeroThat(iptr *int) {
	*iptr = 0
}

func changeName(name *string) {
	*name = "Matt"
}

func main() {
	initial := 27
	fmt.Println("initial", initial)

	zeroThis(initial)
	fmt.Println("zeroThis initial", initial)

	zeroThat(&initial)
	fmt.Println("zeroThis initial", initial)
	fmt.Println("pointer initial", &initial)

	name := "matheus"
	fmt.Println("name", name)
	changeName(&name)
	fmt.Println("name", name)
}
