package context2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Start() {
	http.HandleFunc("/data", getDataHandler)

	// Start the HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getDataHandler handles the GET request and uses a context with a 10-second timeout
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context that will automatically cancel after 10 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel() // Always call cancel to release resources

	// Simulate a long-running process (e.g., database query, external API call)
	result, err := fetchData(ctx)
	if err != nil {
		// If the context is canceled or times out, respond with an error
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	// If everything is fine, return the result
	fmt.Fprintf(w, "Data: %s", result)
}

// fetchData simulates a long-running operation with context support
func fetchData(ctx context.Context) (string, error) {
	// Simulate a process that might take 12 seconds to complete
	select {
	case <-time.After(12 * time.Second):
		return "Fetched Data Successfully", nil
	case <-ctx.Done():
		// If the context was canceled or timed out, return an error
		return "", fmt.Errorf("request canceled or timed out: %v", ctx.Err())
	}
}
