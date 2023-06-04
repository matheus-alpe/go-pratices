package main

import (
	"fmt"
	"time"
)

func Print(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, s)
	}
}

func main() {
	fmt.Println("Start")
	go Print("First")
	go Print("Second")
	time.Sleep(5 * time.Second)
	fmt.Println("End")
}
