package checker

import "net/http"

// Describes the outcome of a URL check.
type Result struct {
	Url        string
	IsLive     bool
	StatusCode int
	Status     string
	Err        error
}

func CheckURL(url string, client *http.Client) Result {
	// Sends a request and prints the response.
	// Uses Head() over Get() as we have no need for the request body.
	resp, err := client.Head(url)

	// Gives the result the error if we failed to recieve a response.
	if err != nil {
		return Result{Url: url, IsLive: false, Err: err}
	}

	// Closes the response when the function ends.
	// Placed here to avoid an error if a response is not recieved.
	defer resp.Body.Close()
	statusCode := resp.StatusCode
	status := resp.Status

	// Returns the Result
	return Result{Url: url, IsLive: true, StatusCode: statusCode, Status: status}
}
