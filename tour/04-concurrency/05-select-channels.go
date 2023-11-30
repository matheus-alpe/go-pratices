package concurrency

import (
	"fmt"
)

type QuitChannel chan struct{}

func fibonacciSelect(c chan int, quit QuitChannel) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func SelectChannelsExample() {
	fmt.Println("\nSelect Channels:")

	c := make(chan int)
	quit := make(QuitChannel)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		close(quit)
	}()

	fibonacciSelect(c, quit)
}
