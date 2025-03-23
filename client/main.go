package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// APIBaseURL is the target API Gateway URL
var APIBaseURL = getEnv("API_BASE_URL", "http://192.168.1.206:8084/serv0/home")

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

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error request: ", err)
		return
	}

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.VYeo_tt-VouYVxcxRfDVGA7Mwc9h2AYuk_4W5errbEA")


	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	fmt.Println("status code: ", resp.StatusCode)

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("GET Response: %s\n", string(body))
}

func main() {
	fmt.Println("Starting API Client...")

	// Test GET request
	SendGETRequest("/")
}
