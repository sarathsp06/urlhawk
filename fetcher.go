package main

// URLResponse represents the response from a URLHawk fetch operation
type URLResponse struct {
	URL        string `json:"url"`
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
	Error      string `json:"error,omitempty"`
}

// FetchResults - URLHawk's main hunting function
// TODO: Implement URLHawk's concurrent URL fetching capabilities
func FetchResults(urls []string) ([]URLResponse, error) {
	// Your implementation here - make URLHawk soar! 🦅
	//
	// Hints for URLHawk's hunting strategy:
	// - Use net/http package for HTTP requests
	// - Consider using goroutines for concurrent strikes (hawk-like speed!)
	// - Handle errors gracefully (even hawks miss sometimes)
	// - Read response body and close it properly (clean hunting)
	// - Set appropriate timeouts (hawks don't wait forever)

	return make([]URLResponse, len(urls)), nil
}
