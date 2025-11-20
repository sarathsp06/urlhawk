package main

import (
	"testing"
)

// TestFetchResults tests the basic functionality
// FIX: Fix the test
func TestFetchResults(t *testing.T) {
	// Test with valid URLs
	urls := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/json",
	}

	results, err := FetchResults(urls)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(results) != len(urls) {
		t.Fatalf("Expected %d results, got %d", len(urls), len(results))
	}
}
