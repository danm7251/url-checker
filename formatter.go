package main

import (
	"fmt"
	"strings"
)

// Holds formatting data.
type Formatter struct {
	urlPadding int
	rawErrors  bool
}

func NewFormatter(urls []string, rawErrors bool) Formatter {
	return Formatter{urlPadding: getMaxWidth(urls) + 1, rawErrors: rawErrors}
}

func (f Formatter) printResult(result Result) {
	// ANSI colour coded keywords.
	const (
		Up   = "\033[1;32mUP\033[0m"
		Down = "\033[1;31mDOWN\033[0m"
	)

	if result.isLive {
		fmt.Printf("%-*s | %s   | %s\n", f.urlPadding, result.url, Up, result.status)
	} else {
		var errMsg string

		// Translates errors to a more compact format if feature is not disabled.
		if f.rawErrors {
			errMsg = result.err.Error()
		} else {
			errMsg = translateError(result.err)
		}

		fmt.Printf("%-*s | %s | %v\n", f.urlPadding, result.url, Down, errMsg)
	}
}

// Finds the length of the longest string.
func getMaxWidth(urls []string) int {
	maxWidth := 0

	for _, url := range urls {
		if len(url) > maxWidth {
			maxWidth = len(url)
		}
	}

	return maxWidth
}

// Matches errors to provide shorthand formats.
func translateError(err error) string {
	msg := err.Error()

	switch {
	case strings.Contains(msg, "refused"):
		return "CONNECTION REFUSED"
	case strings.Contains(msg, "no such host"):
		return "NO SUCH HOST"
	case strings.Contains(msg, "certificate has expired"):
		return "SSL: EXPIRED/INVALID"
	case strings.Contains(msg, "signed by unknown authority"):
		return "SSL: UNKNOWN AUTHORITY"
	case strings.Contains(msg, "certificate is valid for"):
		return "SSL: HOSTNAME MISMATCH"
	case strings.Contains(msg, "Client.Timeout"):
		return "TIMEOUT"
	default:
		return msg
	}
}
