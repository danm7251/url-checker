package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Creates an HTTP client with a 5-second timeout.
	client := &http.Client{Timeout: 5 * time.Second}

	for _, url := range os.Args[1:] {
		// Adds an https:// prefix if the URL has no protocol.
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		// Sends a request and prints the response.
		resp, err := client.Get(url)

		// Prints the error if we failed to recieve a response.
		if err != nil {
			fmt.Printf("%v: %v\n", url, err)
			continue
		}

		// Prints the status code and closes the response.
		fmt.Printf("%v: %v\n", url, resp.StatusCode)
		resp.Body.Close()
	}
}
