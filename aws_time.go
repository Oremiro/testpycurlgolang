package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
	"time"
)

func main() {
	url := "https://aws.okx.com/api/v5/public/time"
	timeout := 0.5 // seconds

	// Create an HTTP client with HTTP/2 support
	client := &http.Client{
		Transport: &http2.Transport{},
	}

	var totalDuration time.Duration

	for i := 1; ; i++ {
		startTime := time.Now()

		// Make the request
		response, err := client.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = response.Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
		elapsedTime := time.Since(startTime)

		// Calculate and print average time after each iteration
		if i != 1 {
			totalDuration += elapsedTime
			averageDuration := totalDuration / time.Duration(i-1)
			fmt.Printf("[GO] Average time after %d requests: %s\n\n", i-1, averageDuration)
		}

		time.Sleep(time.Duration(timeout * float64(time.Second)))
	}
}
