package methodsinterfaces

import "fmt"

// An empty interface may hold values of any type. (Every type implements at least zero methods.)
func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func InterfacesEmptyExample() {
	fmt.Println("\nInterfaces Empty:")

	var i interface{}
	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "hello"
	describeEmpty(i)
}
