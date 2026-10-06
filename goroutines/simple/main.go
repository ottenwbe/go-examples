package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	start := time.Now()

	go func() {
		defer wg.Done()
		for i := range 100 {
			fmt.Printf("I am a goroutine (%d/100).\n", i+1)
			// The sleep only simulates work. Never use time.Sleep to
			// synchronize goroutines; that is what the WaitGroup below is for.
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Printf("Goroutine finished after %v.\n", time.Since(start).Round(time.Millisecond))
}
