package basics

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func exampleFuncClosure1() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func FunctionClosuresExample() {
	fmt.Println("\nFunctions closures:")

	exampleFuncClosure1()
}
