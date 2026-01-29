package main

import (
	"net/http"
)

func main() {
	// Parses the flags and creates the config.
	config := LoadConfig()

	// Parses the URLs from the arguments.
	urls, err := ParseArgs()
	// The only error is already handled internally.
	if err != nil {
		return
	}

	// Creates an HTTP client with the desired timeout.
	client := &http.Client{Timeout: config.Timeout}

	// Creates a buffered channel to collect Results.
	results := make(chan Result, len(urls))

	// Loops through all URLs provided.
	for _, url := range urls {
		// An anonymous goroutine that runs checkURL()
		go func() {
			// Sends the Result to the channel.
			results <- checkURL(url, client)
		}()
	}

	// Creates a new formatter to handle output.
	formatter := NewFormatter(urls, config.RawErrors)

	// Takes each available result from the channel and prints them.
	// Blocks the main loop until all URL results are taken.
	for range len(urls) {
		result := <-results
		formatter.printResult(result)
	}

	// Closes write access to the channel.
	close(results)
}

// Describes the outcome of a URL check.
type Result struct {
	url        string
	isLive     bool
	statusCode int
	status     string
	err        error
}

func checkURL(url string, client *http.Client) Result {
	// Sends a request and prints the response.
	// Uses Head() over Get() as we have no need for the request body.
	resp, err := client.Head(url)

	// Gives the result the error if we failed to recieve a response.
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
