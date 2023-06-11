package basics

import (
	"fmt"
	// "golang.org/x/tour/wc"
)

// https://go.dev/tour/moretypes/23
func fibonacci() func() int {
	var n1 = 0
	var n2 = 1

	return func() int {
		n := n1
		n1 = n2
		n2 = n + n2
		return n
	}
}

func PrintExerciseClosure() {
	fmt.Println("\nExercise 03: Maps")

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
