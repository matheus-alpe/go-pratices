package basics

import "fmt"

func forExample() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func forContinuedExample() {
	sum := 1

	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)
}

func forWhileExample() {
	sum := 1

	for sum < 200 {
		sum += sum
	}

	fmt.Println(sum)
}

/* Loop forever
for {} // this loop will spin, using 100% CPU
*/

func PrintForExemple() {
	fmt.Println("\nFor:")

	forExample()
	forContinuedExample()
	forWhileExample()
}
