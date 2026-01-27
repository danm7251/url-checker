package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// Adds an https:// prefix if the URL has no protocol
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		// Send request and print response
		response, err := http.Get(url)
		fmt.Printf("\n%v\nResponse: %v\nError: %v\n", url, response, err)
	}
}
