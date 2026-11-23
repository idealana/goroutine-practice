// LEVEL 1 - CHAPTER 4 - Buffered and Unbuffered Channel

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	runChannel("Buffered Channel", make(chan int, 5))
	runChannel("Unbuffered Channel", make(chan int))
}

func runChannel(name string, ch chan int) {
	timeStart := time.Now()

    var producerTime time.Duration
    var consumerTime time.Duration
    
	var wg sync.WaitGroup
    wg.Add(2)
    
    // producer
    go func(ch chan<- int) {
    	defer wg.Done()
        
        start := time.Now()
               
    	numbers := []int{1,2,3,4,5,6,7,8,9,10}
        
        for i, val := range numbers {
        	ch <- val

        	if i < len(numbers) - 1 {
        		time.Sleep(100 * time.Millisecond)
        	}
        }
        
        close(ch)
        
        producerTime = time.Since(start)
    }(ch)
    
    // consumer
    go func(ch <-chan int) {
    	defer wg.Done()
        
        start := time.Now()
        
    	for range ch {
        	time.Sleep(500 * time.Millisecond)
        }
        
        consumerTime = time.Since(start)
    }(ch)
    
    wg.Wait()
    
    fmt.Printf("--- %s ---\n", name)
    fmt.Println("Producer completed:", producerTime)
    fmt.Println("Consumer completed:", consumerTime)
    fmt.Println("Total Execution:", time.Since(timeStart))
}
