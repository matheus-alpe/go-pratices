package basics

import (
	"fmt"
	"math"
)

func printTypeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func printTypeInference() {
	v := 42
	f := 3.142
	g := 0.867 + 0.5i

	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}

func PrintTypeConversionExample() {
	fmt.Println("\nType Conversion:")

	printTypeConversion()
	printTypeInference()
}
