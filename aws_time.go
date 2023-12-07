package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://aws.okx.com/api/v5/public/time"
	numRequests := 201

	client := &http.Client{}

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

	averageDuration := totalDuration / time.Duration(numRequests-1)
	fmt.Printf("\ngolang: Average time for %d requests: %s\n", numRequests, averageDuration)
}
