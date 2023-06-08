package basics

import "fmt"

func ForExample() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func ForContinuedExample() {
	sum := 1

	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)
}

func ForWhileExample() {
	sum := 1

	for sum < 200 {
		sum += sum
	}

	fmt.Println(sum)
}

/* Loop forever
for {} // this loop will spin, using 100% CPU
*/
