// LEVEL 1 - CHAPTER 1 - Number Printer

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	totalGoroutines := 3

	wg.Add(totalGoroutines)

	for i := 1; i <= totalGoroutines; i++ {
		start := (i - 1) * 10 + 1
		end := i * 10

		go func(id, start, end int) {
			defer wg.Done()

			for number := start; number <= end; number++ {
				// fmt.Println("Goroutine", id, ":", number)
				fmt.Printf("Goroutine %d: %d\n", id, number)

				time.Sleep(500 * time.Millisecond)
			}
		}(i, start, end)
	}

	// go func() {
	// 	defer wg.Done()

	// 	for i := 1; i <= 10; i++ {
	// 		fmt.Println("Goroutine 1:", i)

	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()

	// 	for i := 11; i <= 20; i++ {
	// 		fmt.Println("Goroutine 2:", i)

	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()

	// 	for i := 21; i <= 30; i++ {
	// 		fmt.Println("Goroutine 3:", i)

	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }()

	wg.Wait()

	fmt.Println("Goroutines completed")
}
