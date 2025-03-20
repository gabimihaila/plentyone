package main

import (
	"fmt"
	"net/http"
	"os"
)

// Handler function for HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! I'm a microservice!")
}


// Function to get environment variables with default values
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "8081")

	http.HandleFunc("/", handler)

	fmt.Println("Microservice is running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
