# TODO List
## Crucial
- [x] Print results live.
- [x] Cover error cases.
- [x] Input sanitisation e.g whitespace.
- [x] Add usage example in case of help flag or no arguments.

## Potential
- [ ] Decide whether to focus on reachability or behaviour. e.g in the case of expired certificates.
- [x] Passing files as arguments.
- [ ] Outputting to files.
- [ ] Create tests.

## Issues
- [ ] Limit concurrent behaviour in case of many URLs
- [ ] Limit memory usage for large files with bufio
- [ ] Solidify error handling (Errors in cli.go don't print)