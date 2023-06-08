package basics

import "fmt"

var c, python, java bool
var j, k int = 1, 2

// message := "message" // error: needs to be var message = "message"

const Pi = 3.14

func PrintVars() {
	var i int
	fmt.Println(i, c, python, java)

	var golang, rust, typescript = true, true, "no"
	fmt.Println(k, j, golang, rust, typescript)

	l := "lang"
	fmt.Println(l)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func NeedInt(x int) int {
	return x*10 + 1
}

func NeedFloat(x float64) float64 {
	return x * 0.1
}
