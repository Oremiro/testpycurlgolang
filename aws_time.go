package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
	"time"
)

func main() {
	url := "https://aws.okx.com/api/v5/public/time"
	numRequests := 201

	client := &http.Client{
		Transport: &http2.Transport{},
	}

	var totalDuration time.Duration

	for i := 1; i <= numRequests; i++ {
		startTime := time.Now()

		response, err := client.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer response.Body.Close()

		elapsedTime := time.Since(startTime)

		if i != 1 {
			totalDuration += elapsedTime
		}

		fmt.Printf("Attempt %d: HTTP GET request to %s completed in %s\n", i, url, elapsedTime)
	}

	averageDuration := totalDuration / time.Duration(numRequests)
	fmt.Printf("\nAverage time for %d requests: %s\n", numRequests, averageDuration)
}
