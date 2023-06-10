package basics

import "fmt"

// (x, y int) = (x int, y int)
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// named returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

func PrintFunctionsExample() {
	fmt.Println("\nFunctions:")

	fmt.Println(add(2, 5))
	fmt.Println(swap("Matheus", "Thiago"))
	fmt.Println(split(20))
}
