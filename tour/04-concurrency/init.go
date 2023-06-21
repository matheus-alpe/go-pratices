package concurrency

import "fmt"

func init() {
	fmt.Println("\nConcurrency:")

	GoRoutinesExample()
	ChannelsExample()
	BufferedChannelsExample()
}
