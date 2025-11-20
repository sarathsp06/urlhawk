package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Parse command line flags
	outputFile := flag.String("output", "results.csv", "Output file for results")
	help := flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *help {
		printUsage()
		return
	}

	// Read URLs from stdin
	urls, err := readURLsFromStdin()
	if err != nil {
		log.Fatal("Error reading URLs:", err)
	}

	if len(urls) == 0 {
		fmt.Println("No URLs provided. Please provide URLs via stdin.")
		printUsage()
		return
	}

	fmt.Printf("Fetching %d URLs...\n", len(urls))

	// Fetch results
	results, err := FetchResults(urls)
	if err != nil {
		log.Fatal("Error fetching URLs:", err)
	}

	// Write results to file
	file, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer file.Close()

	// Write CSV header
	fmt.Fprintln(file, "URL,StatusCode,BodyLength,Error")

	// Write each result
	for _, result := range results {
		err = writeResultsToFile(result, file)
		if err != nil {
			log.Fatal("Error writing result:", err)
		}
	}

	fmt.Printf("Results written to %s\n", *outputFile)
}

func readURLsFromStdin() ([]string, error) {
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	return urls, scanner.Err()
}

func writeResultsToFile(result URLResponse, file *os.File) error {
	// TODO: Implement CSV row writing functionality
	return nil
}

func printUsage() {
	fmt.Println("URLHawk - Fast URL Fetcher CLI")
	fmt.Println("Usage: echo 'https://example.com' | go run . [flags]")
	fmt.Println("   or: go run . [flags] < urls.txt")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  echo 'https://httpbin.org/json' | go run .")
	fmt.Println("  echo -e 'https://httpbin.org/status/200\\nhttps://httpbin.org/json' | go run . -output=my_results.csv")
	fmt.Println("  go run . < urls.txt")
}
