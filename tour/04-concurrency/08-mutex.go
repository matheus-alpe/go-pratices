package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	wg sync.WaitGroup
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
	c.wg.Done()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func MutexExample() {
	fmt.Println("\nMutex:")
	start := time.Now()

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1e6; i++ {
		c.wg.Add(1)
		go c.Inc("somekey")
	}
	c.wg.Wait()

	fmt.Println(c.Value("somekey"), time.Since(start))
}
