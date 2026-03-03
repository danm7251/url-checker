package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Options struct {
	Timeout   time.Duration
	RawErrors bool
}

func Load() (Options, []string, error) {
	var opts Options
	var fromFile bool

	flag.DurationVar(&opts.Timeout, "t", 5*time.Second, "")
	flag.DurationVar(&opts.Timeout, "timeout", 5*time.Second, "")
	flag.BoolVar(&opts.RawErrors, "e", false, "")
	flag.BoolVar(&opts.RawErrors, "errors", false, "")
	flag.BoolVar(&fromFile, "f", false, "")
	flag.BoolVar(&fromFile, "file", false, "")

	flag.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: ./urlcheck [OPTIONS] [URL 1] [URL 2] ...")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("\t-h, --help\tDisplay help")
		fmt.Println("\t-t, --timeout\tSet request timeout (default=5s)")
		fmt.Println("\t-e, --errors\tShow unformatted error messages")
		fmt.Println("\t-f, --file\tRead URLs from a file:\n\tUsage: ./urlcheck [OPTIONS] -f [FILENAME]")
		fmt.Println()
		fmt.Println("URLs:")
		fmt.Println("\tArguments should be whitespace seperated URLs.")
		fmt.Println("\tMissing protocols will default to HTTPS.")
		fmt.Println()
	}

	flag.Parse()

	// Parse URLs
	if fromFile {
		if len(flag.Args()) != 1 {
			flag.Usage()
			return Options{}, nil, fmt.Errorf("exactly one filename must be provided")
		}

		urls, err := readURLsFromFile(flag.Arg(0))
		if err != nil {
			return Options{}, nil, fmt.Errorf("could not read file: %v", err)
		}

		return opts, sanitiseURLs(urls), nil
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
		return Options{}, nil, fmt.Errorf("no arguments provided")
	}

	return opts, sanitiseURLs(flag.Args()), nil
}

func readURLsFromFile(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	urls := strings.Split(string(bytes), "\n")

	return urls, nil
}

func sanitiseURLs(urls []string) []string {
	for i, url := range urls {
		// Removes any whitespace
		url = strings.TrimSpace(url)

		// Discards empty strings
		if url == "" {
			continue
		}

		// Adds an https:// prefix if the URL has no protocol.
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		urls[i] = url
	}

	return urls
}
