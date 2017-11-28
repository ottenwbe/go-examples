package main

import (
	"fmt"
	"time"
)

func main() {

	results := make(chan bool)

	go func(res chan bool) {
		for i := 0; i < 100; i++ {
			fmt.Printf("I am a goroutine (%d/100).\n", i+1)
			time.Sleep(100)
		}
		res <- true
	}(results)

	_ = <-results
}
