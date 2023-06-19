package concurrency

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func GoRoutinesExample() {
	fmt.Println("\nGo Routines:")

	go say("World")
	say("Hello")
}
