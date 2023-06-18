package basics

import "fmt"

var c, python, java bool
var j, k int = 1, 2

// message := "message" // error: needs to be var message = "message"

const pi = 3.14

func printVars() {
	var i int
	fmt.Println(i, c, python, java)

	var golang, rust, typescript = true, true, "no"
	fmt.Println(k, j, golang, rust, typescript)

	l := "lang"
	fmt.Println(l)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func PrintVarsExample() {
	fmt.Println("\nVars & Constants:")

	printVars()
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
