package methodsinterfaces

import "fmt"

type RandomType int

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case RandomType:
		fmt.Printf("%v is custom type of: %T\n", v, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeSwitchesExample() {
	fmt.Println("\nType Switches:")

	do(21)
	do("Hello")
	do(RandomType(23))
	do(true)
}
