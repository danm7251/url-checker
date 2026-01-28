* Development Log
**Jan 27, 2026 - 15:30**
> * Initialised the module.
> * Created README file.
> * Initialised git.

**Jan 27, 2026 - 15:45**
> * Read about os.Args and implemented it.
> * Discovered that range {array} returns the index and then the value.
> * Discovered that Go has _ for ignoring values.
> * Wrote code for the main function that reads the arguments, loops through and prints them

**Jan 27, 2026 - 16:00**
> * Imported the net/http package.
> * Used http.Get() to check URLs.
> * Discovered http.Get() requires the protocol to be provided.
> * Wrote code that checks for a http prefix and adds one if missing.
> * Main now loops through URLs, checks them and prints the response.

**Jan 27, 2026 - 16:55**
> * Returned home and picked up where I left off.
> * Created a repository on Github and uploaded this project.
> * Minor fix to the README.
> * Created an HTTP client with the net package.
> * Imported the time package to set the clients timeout to 5 seconds.

**Jan 27, 2026 - 17:15**
> * Separate logic branch for failed requests.
> * Moved the HTTP client out of the for loop and used & to reference it avoiding copying.
> * Closed the response at the end of the loop.

**Jan 27, 2026 - 17:30**
> * Moved the logic into a separate function.
> * Discovered defer, learned it will be useful for when I introduce concurrency.
> * Used defer resp.Body.Close() after the error check to avoid trying to close a non-existent response.
> * Used ANSI colour codes to colour the status code green when it is 200.

**Jan 28, 2026 - 17:00**
> * Had a slight setback where my Laptop broke yesterday evening. It is now up and running.
> * Tried using the go keyword to check URLs concurrently.
> * Realised they weren't printing output as the main function was terminating the program before they could finish.
> * Used time.Sleep() however it's clearly not the answer since the program waits even after the URLs are checked.
> * Learning about sync.WaitGroup to block execution until all goroutines return.

**Jan 28, 2026 - 17:15**
> * Implemented a WaitGroup and now my URL checks are concurrent!

**Jan 28, 2026 - 17:45**
> * Discovered anonymous goroutines.
> * Used an anonymous goroutine to move WaitGroup logic out of checkURL().

**Jan 28, 2026 - 18:00**
> * Learnt about structs.
> * Defined a Result struct containing URL check information.
> * Changed checkURL to return a Result and removed all printing from it.

**Jan 28, 2026 - 18:15**
> * Learnt about channels.
> * Made a channel to collect Results from goroutines.
> * Encountered an issue where main() never terminates.

**Jan 28, 2026 - 20:00**
> * Discovered that the goroutines were blocking.
> * Learnt about buffered vs unbuffered channels.
> * Implemented a buffered channel instead to solve the issue.
> * Learnt that closing channels only removes write access.
> * Moved close(results) to just after the goroutines.

**Jan 28, 2026 - 20:15**
> * Started writing a function dedicated to printing the results.
> * Discovered that as channels are streams they are not indexed.
> * Added UP and DOWN status to the output coloured with ANSI escape codes.

**Jan 28, 2026 - 21:00**
> * Added rudimentary checks to shorten common error messages.
> * Used dynamic padding to neaten the display.