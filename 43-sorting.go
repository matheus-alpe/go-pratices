package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"c", "a", "d", "b"}
	sort.Strings(strings)
	fmt.Println("String", strings)

	ints := []int{7, 1, 5, 2, 4, 3, 6}
	sort.Ints(ints)
	fmt.Println("Ints", ints)

	isIntsSorted := sort.IntsAreSorted(ints)
	fmt.Println("Ints are sorted?", isIntsSorted)
}
