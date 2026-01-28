package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// Creates an HTTP client with a 5-second timeout.
	client := &http.Client{Timeout: 5 * time.Second}

	// Creates a WaitGroup to manage goroutines.
	var wg sync.WaitGroup

	// Creates a channel to collect Results.
	results := make(chan Result)

	// Loops through all URLs provided.
	for _, url := range os.Args[1:] {
		// Increments the WaitGroup counter.
		wg.Add(1)

		// An anonymous goroutine that runs checkURL()
		// and decrements the WaitGroup counter when it returns.
		go func() {
			defer wg.Done()
			// Sends the Result to the channel.
			results <- checkURL(url, client)
		}()
	}

	// Blocks until all goroutines return.
	wg.Wait()

	// Outputs the results.
	fmt.Printf("%v", results)

	// Closes the channel once we have finished using it.
	close(results)
}

type Result struct {
	url        string
	isLive     bool
	statusCode int
	err        error
}

func checkURL(url string, client *http.Client) Result {
	// Adds an https:// prefix if the URL has no protocol.
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	// Sends a request and prints the response.
	resp, err := client.Get(url)

	// Prints the error if we failed to recieve a response.
	if err != nil {
		return Result{url: url, isLive: false, err: err}
	}

	// Closes the response when the function ends.
	// Placed here to avoid an error if a response is not recieved.
	defer resp.Body.Close()
	statusCode := resp.StatusCode

	// Returns the Result
	return Result{url: url, isLive: true, statusCode: statusCode}
}
