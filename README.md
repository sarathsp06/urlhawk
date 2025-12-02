# URLHawk - CLI Interview Challenge

A CLI tool for fetching HTTP responses from multiple URLs concurrently and writing results to CSV format.

## Problem Statement

Implement a CLI application that:

1. Reads URLs from stdin (one per line)
2. Fetches HTTP GET responses from those URLs concurrently
3. Writes the results to a CSV file


## Requirements

1. The function should accept a slice of URL strings
2. Make HTTP GET requests to each URL
3. Return the responses in a structured format
4. Handle errors appropriately
5. Consider concurrent execution for better performance
6. CLI should read URLs from stdin and write results to CSV file

## Getting Started

1. Implement the `FetchResults` function in `fetcher.go`
2. Implement the `writeResultsToFile` function in `main.go` for CSV output
3. Run tests: `go test`
4. Test the CLI: `echo 'https://httpbin.org/json' | go run .`
5. **Important**: Fix the deliberately flawed tests after implementing the functions

## Test Debugging Challenge

The provided test file contains intentional bugs. After implementing your functions:

1. Run `go test -v` to see test results
2. Analyze the test code to identify logical flaws
3. Fix the tests to properly validate your implementation
4. Ensure all tests pass with correct implementation

## CLI Usage

```bash
# Single URL
echo 'https://httpbin.org/json' | go run .

# Multiple URLs
echo -e 'https://httpbin.org/status/200\nhttps://httpbin.org/json' | go run .

# From file
go run . < urls.txt

# Custom output file
echo 'https://httpbin.org/json' | go run . -output=my_results.csv

# Show help
go run . -help
```
