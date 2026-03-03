package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	Timeout   time.Duration
	RawErrors bool
	FromFile  bool
	Args      []string
}

func LoadConfig() Config {
	var timeout time.Duration
	var rawErrors bool
	var fromFile bool

	flag.DurationVar(&timeout, "t", 5*time.Second, "")
	flag.DurationVar(&timeout, "timeout", 5*time.Second, "")
	flag.BoolVar(&rawErrors, "e", false, "")
	flag.BoolVar(&rawErrors, "errors", false, "")
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

	return Config{Timeout: timeout, RawErrors: rawErrors, FromFile: fromFile, Args: flag.Args()}
}

// Assumes that the flags have already been parsed.
func ParseArgs(config Config) ([]string, error) {
	if config.FromFile {
		if len(config.Args) != 1 {
			flag.Usage()
			return nil, fmt.Errorf("exactly one filename must be provided")
		}

		urls, err := readURLsFromFile(config.Args[0])
		if err != nil {
			return nil, fmt.Errorf("could not read file: %s", err)
		}

		return sanitiseURLs(urls), nil
	}

	if len(config.Args) == 0 {
		flag.Usage()
		return nil, fmt.Errorf("no arguments provided")
	}

	return sanitiseURLs(config.Args), nil
}

func readURLsFromFile(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	urls := strings.Split(string(bytes), " ")

	return urls, nil
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
