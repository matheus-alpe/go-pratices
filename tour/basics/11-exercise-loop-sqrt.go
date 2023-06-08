package basics

import (
	"fmt"
	"math"
)

func LoopFindSqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func PrintExerciseLoopSqrt() {
	fmt.Println(LoopFindSqrt(234))
	fmt.Println(math.Sqrt(234))
}
