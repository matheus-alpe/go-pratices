package basics

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(4, 5)
}

func exampleFuncValues1() {
	hypo := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("hypo", hypo(5, 12))

	fmt.Println("compute hypo", compute(hypo))
	fmt.Println("compute math.Pow", compute(math.Pow))
}

func FunctionValuesExample() {
	fmt.Println("\nFunctions values:")

	exampleFuncValues1()
}
