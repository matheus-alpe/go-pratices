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

func deferMutation() (i int) {
	defer func() { i++ }()
	return 1
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func DeferExample() {
	fmt.Println("\nDefer:")
	fmt.Println("a")
	defer fmt.Println("b")
	deferPrint()
	deferStack()
	fmt.Println("\nc")
	fmt.Println("deferMutation value:", deferMutation())
	f()
	fmt.Println("Returned normally from f.")
}
