package main

import (
	"net/http"

	"github.com/danm7251/url-checker/internal/checker"
	"github.com/danm7251/url-checker/internal/cli"
	"github.com/danm7251/url-checker/internal/formatter"
)

func main() {
	// Parses the flags and creates the config.
	config := cli.LoadConfig()

	// Parses the URLs from the arguments.
	urls, err := cli.ParseArgs(config)
	// The only error is already handled internally.
	if err != nil {
		return
	}

	// Creates an HTTP client with the desired timeout.
	client := &http.Client{Timeout: config.Timeout}

	// Creates a buffered channel to collect Results.
	results := make(chan checker.Result, len(urls))

	// Loops through all URLs provided.
	for _, url := range urls {
		// An anonymous goroutine that runs checkURL()
		go func() {
			// Sends the Result to the channel.
			results <- checker.CheckURL(url, client)
		}()
	}

	// Creates a new formatter to handle output.
	f := formatter.NewFormatter(urls, config.RawErrors)

	// Takes each available result from the channel and prints them.
	// Blocks the main loop until all URL results are taken.
	for range len(urls) {
		result := <-results
		f.PrintResult(result)
	}

	// Closes write access to the channel.
	close(results)
}
