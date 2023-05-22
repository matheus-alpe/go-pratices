package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println("i:", i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("j:", j)
	}

	stop := false
	for {
		fmt.Println("Loop")
		if stop {
			break
		}
		stop = true
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println("n:", n)
	}
}
