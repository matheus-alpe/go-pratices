package tour

import "fmt"

var c, python, java bool
var j, k int = 1, 2

// message := "message" // error: needs to be var message = "message"

func PrintVars() {
	var i int
	fmt.Println(i, c, python, java)

	var golang, rust, typescript = true, true, "no"
	fmt.Println(k, j, golang, rust, typescript)

	l := "lang"
	fmt.Println(l)
}
