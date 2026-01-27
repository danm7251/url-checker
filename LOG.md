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