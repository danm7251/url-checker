package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		// Adds an https:// prefix if the URL has no protocol.
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		// Creates an HTTP client with a 5-second timeout.
		client := http.Client{Timeout: 5 * time.Second}

		// Sends a request and prints the response.
		resp, err := client.Get(url)

		if err != nil {
			fmt.Printf("%v: %v\n", url, err)
			continue
		}

		fmt.Printf("%v: %v\n", url, resp.StatusCode)
	}
}
