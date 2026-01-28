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
	// Creates a slice of the arguments containing the URLs
	urls := os.Args[1:]

	// Finds the longest URL for formatting
	urlWidth := 0
	for _, url := range urls {
		if len(url) > urlWidth {
			urlWidth = len(url)
		}
	}

	// Creates an HTTP client with a 5-second timeout.
	client := &http.Client{Timeout: 5 * time.Second}

	// Creates a WaitGroup to manage goroutines.
	var wg sync.WaitGroup

	// Creates a buffered channel to collect Results.
	results := make(chan Result, len(urls))

	// Loops through all URLs provided.
	for _, url := range urls {
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

	// Blocks until all goroutines return and closes write access to the channel.
	wg.Wait()
	close(results)

	// Outputs the results.
	printResults(results, urlWidth)
}

type Result struct {
	url        string
	isLive     bool
	statusCode int
	status     string
	err        error
}

func checkURL(url string, client *http.Client) Result {
	// Adds an https:// prefix if the URL has no protocol.
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	// Sends a request and prints the response.
	// Uses Head() over Get() as we have no need for the request body.
	resp, err := client.Head(url)

	// Prints the error if we failed to recieve a response.
	if err != nil {
		return Result{url: url, isLive: false, err: err}
	}

	// Closes the response when the function ends.
	// Placed here to avoid an error if a response is not recieved.
	defer resp.Body.Close()
	statusCode := resp.StatusCode
	status := resp.Status

	// Returns the Result
	return Result{url: url, isLive: true, statusCode: statusCode, status: status}
}

func printResults(results chan Result, urlWidth int) {
	const (
		Up   = "\033[1;32mUP\033[0m"
		Down = "\033[1;31mDOWN\033[0m"
	)

	// Ensures the longest URL lines up in the display.
	urlWidth = urlWidth + 1

	for result := range results {
		if result.isLive {
			fmt.Printf("%-*s | %s   | %s\n", urlWidth, result.url, Up, result.status)
		} else {
			err := translateError(result.err)
			fmt.Printf("%-*s | %s | %v\n", urlWidth, result.url, Down, err)
		}
	}
}

func translateError(err error) string {
	msg := err.Error()

	switch {
	case strings.Contains(msg, "refused"):
		return "CONNECTION REFUSED"
	case strings.Contains(msg, "no such host"):
		return "NO SUCH HOST"
	default:
		return msg
	}
}
