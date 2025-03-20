package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// APIBaseURL is the target API Gateway URL
var APIBaseURL = getEnv("API_BASE_URL", "http://192.168.1.137:8081")

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// SendGETRequest sends a GET request to the API
func SendGETRequest(endpoint string) {
	url := APIBaseURL + endpoint
	client := &http.Client{Timeout: 10 * time.Second}

	fmt.Println("timeout?: ", client.Timeout)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("GET Response: %s\n", string(body))
}

func main() {
	fmt.Println("Starting API Client...")

	// Test GET request
	SendGETRequest("/")
}
