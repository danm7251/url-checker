package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Timeout   time.Duration
	RawErrors bool
}

func LoadConfig() Config {
	var timeout time.Duration
	var rawErrors bool

	flag.DurationVar(&timeout, "t", 5*time.Second, "")
	flag.DurationVar(&timeout, "timeout", 5*time.Second, "")
	flag.BoolVar(&rawErrors, "e", false, "")
	flag.BoolVar(&rawErrors, "errors", false, "")

	flag.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: ./url-checker [OPTIONS] [URL 1] [URL 2] ...")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("\t-h, --help\tDisplay help")
		fmt.Println("\t-t, --timeout\tSet request timeout (default=5s)")
		fmt.Println("\t-e, --errors\tShow unformatted error messages")
		fmt.Println()
		fmt.Println("URLs:")
		fmt.Println("\tArguments should be whitespace seperated URLs.")
		fmt.Println("\tMissing protocols will default to HTTPS.")
		fmt.Println()
	}

	flag.Parse()

	return Config{Timeout: timeout, RawErrors: rawErrors}
}

func ParseArgs() ([]string, error) {
	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		return nil, fmt.Errorf("no arguments provided")
	}

	return sanitiseURLs(args), nil
}

func sanitiseURLs(urls []string) []string {
	for i, url := range urls {
		// Removes any whitespace
		url = strings.TrimSpace(url)

		// Adds an https:// prefix if the URL has no protocol.
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		urls[i] = url
	}

	return urls
}
