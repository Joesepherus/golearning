package context3

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Data structure for JSON response
type Data struct {
	Message string `json:"message"`
}

func loadPageHandler(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlContent, err := ioutil.ReadFile("./context3/static/index.html")
	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return
	}

	// Write the HTML content to the response writer
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

func Start() {
	http.HandleFunc("/", loadPageHandler)

	// Handle the "/data" endpoint
	http.HandleFunc("/data", getDataHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handles the "/data" request
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	// Context with a 10-second timeout
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := fetchData(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	// Set response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func fetchData(ctx context.Context) (*Data, error) {
	// Simulate a long-running process
	select {
	case <-time.After(5 * time.Second): // Simulate fetching data
		return &Data{Message: "Fetched data successfully!"}, nil
	case <-ctx.Done(): // Handle context cancellation
		return nil, fmt.Errorf("request canceled by user: %v", ctx.Err())
	}
}
