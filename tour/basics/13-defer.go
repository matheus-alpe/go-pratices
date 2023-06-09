package basics

import "fmt"

func deferPrint() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}

	fmt.Println("done")
}

func DeferExample() {
	fmt.Println("a")
	defer fmt.Println("b")
	deferPrint()
	deferStack()
	fmt.Println("\nc")
}
