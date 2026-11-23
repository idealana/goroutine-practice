// LEVEL 2 - CHAPTER 5 - Concurrent Counter

package main

import (
	"fmt"
	"sync"
)

func main() {
	const (
		totalGoroutines = 100
		incrementPerGoroutine = 1000
	)

	counter := 0

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	wg.Add(totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			defer wg.Done()

			for num := 0; num < incrementPerGoroutine; num++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Counter :", counter)
	fmt.Println("Expected :", totalGoroutines * incrementPerGoroutine)
}
