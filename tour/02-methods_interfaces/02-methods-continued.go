package methodsinterfaces

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MethodsContinuedExample() {
	fmt.Println("\nMethods Continued:")

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
