package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println("messages:", <-messages)
	fmt.Println("messages:", <-messages)
}
