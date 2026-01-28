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
	wg := new(sync.WaitGroup)

	// Loops through all URLs provided.
	for _, url := range os.Args[1:] {
		// Increments the WaitGroup counter.
		wg.Add(1)

		// An anonymous goroutine that runs checkURL()
		// and decrements the WaitGroup counter when it returns.
		go func() {
			defer wg.Done()
			checkURL(url, client)
		}()
	}

	// Blocks until all goroutines return.
	wg.Wait()
}

func checkURL(url string, client *http.Client) {
	// ANSI Colour codes
	const (
		ColourGreen = "\033[1;92m"
		ColourReset = "\033[0m"
	)

	// Adds an https:// prefix if the URL has no protocol.
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	// Sends a request and prints the response.
	resp, err := client.Get(url)

	// Prints the error if we failed to recieve a response.
	if err != nil {
		fmt.Printf("%v: %v\n", url, err)
		return
	}

	// Closes the response when the function ends.
	// Placed here to avoid an error if a response is not recieved.
	defer resp.Body.Close()

	// Prints the status code in green if it is 200
	statusCode := resp.StatusCode
	if statusCode == 200 {
		fmt.Printf("%s: %s%d%s\n", url, ColourGreen, statusCode, ColourReset)
		return
	}

	// Prints the status code.
	fmt.Printf("%s: %d\n", url, statusCode)
}
