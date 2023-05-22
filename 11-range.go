package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4, 5}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Println("index:", i, "value:", num)
		}
	}

	fruit_map := map[string]string{"a": "apple", "b": "banana", "l": "lemon"}

	for key, value := range fruit_map {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range fruit_map {
		fmt.Println("key:", key)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
