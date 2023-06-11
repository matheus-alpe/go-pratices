package basics

import (
	"fmt"
	"strings"
)

func example1(primes []int) {
	fmt.Println("\nexample 1:")
	var s []int = primes[2:5]
	fmt.Println(primes)
	fmt.Println(s)
}

func example2() {
	fmt.Println("\nexample 2:")
	names := [4]string{
		"Matheus",
		"Thiago",
		"Marcos",
		"Rosana",
	}

	fmt.Println("names", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("a", a)
	fmt.Println("b", b)

	b[0] = "xxx"
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("names", names)
}

func example3(primes []int) {
	fmt.Println("\nexample 3:")
	r := []bool{true, false, true, true, false, true}
	fmt.Println("r", r)

	type List struct {
		i int
		b bool
	}

	s := []List{}

	for index, value := range primes {
		s = append(s, List{value, r[index]})
	}

	fmt.Println("s", s)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("t", t)
}

func example4(primes []int) {
	fmt.Println("\nexample 4:")
	fmt.Println("primes[1:4]", primes[1:4])

	fmt.Println("primes[:2]", primes[:2])

	fmt.Println("primes[2:]", primes[2:])

	fmt.Println("primes[:]", primes[:])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func example5(primes []int) {
	fmt.Println("\nexample 5:")
	printSlice(primes)
	printSlice(primes[:0])
	printSlice(primes[:4])
	printSlice(primes[2:])
}

func example6() {
	fmt.Println("\nexample 6:")
	w := []int{}
	printSlice(w)

	// if w == nil { //this nil check is never true
	// 	fmt.Println("w nil!")
	// }

	var z []int
	printSlice(z)

	if z == nil {
		fmt.Println("z nil!")
	}
}

func example7() {
	fmt.Println("\nexample 7:")

	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0)
	printSlice(b)

	c := make([]int, 0, 5)
	printSlice(c)

	d := c[:2]
	printSlice(d)

	e := d[2:5]
	printSlice(e)
}

func example8() {
	fmt.Println("\nexample 8:")

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func example9() {
	fmt.Println("\nexample 9:")

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func example10(primes []int) {
	fmt.Println("\nexample 10:")

	for index, value := range primes {
		poweredValue := 1 << uint(value) // == 2**i == math.Pow
		fmt.Printf("[%d] 2**%d = %d\n", index, value, poweredValue)
	}
}

func example11() {
	fmt.Println("\nexample 11:")

	pow := make([]int, 10)

	for index := range pow {
		pow[index] = 1 << uint(index) // == 2**i
	}

	fmt.Println(pow)
	for _, value := range pow {
		fmt.Printf("%d", value)
	}
}

func SliceExample() {
	fmt.Println("\nSlices:")

	numbers := []int{2, 3, 5, 7, 11, 13}
	example1(numbers)
	example2()
	example3(numbers)
	example4(numbers)
	example5(numbers)
	example6()
	example7()
	example8()
	example9()
	example10(numbers)
	example11()
}
