// LEVEL 2 - CHAPTER 2 - Concurrent HTTP Request

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	endpoints := []string{
		"https://jsonplaceholder.typicode.com/users",
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
		"https://jsonplaceholder.typicode.com/todos",
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	wg.Add(len(endpoints))

	execTimeStart := time.Now()

	for _, endpoint := range endpoints {
		go func(endpoint string) {
			defer wg.Done()

			timeStart := time.Now()
			resp, err := client.Get(endpoint)

			if err != nil {
				fmt.Printf(
					"%s request is failed. Error: %v\n",
					endpoint,
					err,
				)
				return
			}

			defer resp.Body.Close()

			_, bodyErr := io.Copy(io.Discard, resp.Body)
			if bodyErr != nil {
				fmt.Printf(
					"%s failed to read response body: %v\n",
					endpoint,
					err,
				)
				return
			}

			timeEnd := time.Since(timeStart)

			statusText := "failed"
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				statusText = "success"
			}

			fmt.Printf(
				"%s request is %s. Status Code: %d, Response Time: %s\n",
				endpoint,
				statusText,
				resp.StatusCode,
				timeEnd.String(),
			)
		}(endpoint)
	}

	wg.Wait()

	execTimeEnd := time.Since(execTimeStart)

	fmt.Printf(
		"Goroutines completed. Execution Time: %s\n",
		execTimeEnd.String(),
	)
}
