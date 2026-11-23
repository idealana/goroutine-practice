// LEVEL 1 - CHAPTER 3 - Channel Communication

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// producer
	go func(ch chan<- int) {
		defer wg.Done()

		numbers := []int{1,2,3,4,5,6,7,8,9,10}

		for _, number := range numbers {
			ch <- number
		}

		close(ch)
	}(ch)

	// consumer
	go func(ch <-chan int) {
		defer wg.Done()

		for number := range ch {
			fmt.Printf("Received: %d, Squared: %d\n", number, number * number)
		}
	}(ch)

	wg.Wait()

	fmt.Println("Goroutines completed")
}
