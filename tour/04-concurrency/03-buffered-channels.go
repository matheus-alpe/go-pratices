package concurrency

import (
	"fmt"
	"time"
)

func BufferedChannelsExample() {
	fmt.Println("\nBuffered Channels:")

	fmt.Println("execution started")
	ch := make(chan int, 2)

	go func() {
		fmt.Println("start goroutine")
		ch <- 1

		time.Sleep(2 * time.Second)

		ch <- 2
		fmt.Println("end goroutine")
	}()

	fmt.Println(<-ch)
	fmt.Println("msg 1")
	// will block this print until finalizes the sleep
	fmt.Println(<-ch)
	fmt.Println("msg 2")

	fmt.Println("execution ended")
}
