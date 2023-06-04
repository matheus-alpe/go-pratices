package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusAll(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	res := plus(1, 4)
	fmt.Println("plus(1, 4) =", res)

	res = plusPlus(1, 4, 5)
	fmt.Println("plusPlus(1, 4, 5) =", res)

	res = plusAll(1, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("plusAll(1, 4, 5, 6, 7, 8, 9, 10) =", res)
}
