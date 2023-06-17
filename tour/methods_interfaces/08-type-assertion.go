package methodsinterfaces

import "fmt"

func TypeAssertionExample() {
	fmt.Println("\nType Assertion:")

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// will panic
	// f = i.(float64)
	// fmt.Println(f)
}
