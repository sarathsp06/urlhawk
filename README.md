# URLHawk - CLI Interview Challenge

**URLHawk** is a high-performance CLI tool that swoops down on multiple URLs simultaneously, fetching their responses with hawk-like precision and speed.

## Problem Statement

Implement a CLI application that:

1. Reads URLs from stdin (one per line)
2. Fetches HTTP GET responses from those URLs concurrently
3. Writes the results to a CSV file

The core challenge is implementing the `FetchResults` function that accepts an array of URLs and returns HTTP GET request responses with optimal performance.

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
3. Run tests with: `go test`
4. Test the CLI with: `echo 'https://httpbin.org/json' | go run .`
5. **IMPORTANT**: The test file contains deliberate bugs! After implementing functions, you must fix the flawed tests.

## Implementation Challenge

You need to implement TWO functions to make URLHawk soar:

### 1. FetchResults Function

- Location: `fetcher.go`
- Purpose: URLHawk's hunting mechanism - fetch HTTP GET responses from multiple URLs
- Should handle errors gracefully and strike with concurrent precision

### 2. writeResultsToFile Function  

- Location: `main.go`
- Purpose: URLHawk's delivery system - write results to CSV nest with proper formatting
- Should include headers: URL, StatusCode, BodyLength, Error

## Function Signature

```go
func FetchResults(urls []string) ([]URLResponse, error)
```

## Expected Output Structure

```go
type URLResponse struct {
    URL        string
    StatusCode int
    Body       string
    Error      string
}
```

## URLHawk CLI Usage

```bash
# Single URL - URLHawk strikes fast! 🦅
echo 'https://httpbin.org/json' | go run .

# Multiple URLs - URLHawk hunts in packs
echo -e 'https://httpbin.org/status/200\nhttps://httpbin.org/json' | go run .

# From file - URLHawk processes your target list
go run . < urls.txt

# Custom output file - URLHawk delivers to your specified nest
echo 'https://httpbin.org/json' | go run . -output=my_results.csv

# Show help - Learn URLHawk's capabilities
go run . -help
```
