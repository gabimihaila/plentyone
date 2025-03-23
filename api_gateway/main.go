package main

import (
	"fmt"
	"log"
	"net/http"

	config "api_gateway/config"
	handler "api_gateway/handlers"
	auth "api_gateway/jwt_auth"
	logger "api_gateway/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	destinations := cfg.Destinations

	mux := http.NewServeMux()
	mux.Handle("/", auth.Auth(http.HandlerFunc(handler.ProxyHandler(destinations))))

	// Initialize logger
	logger.InitLogger()

	// Log ?
	//logger.LoggerMiddleware(http.HandlerFunc(handler.ProxyHandler(destinations, "/serv0")))

	// Start server
	port := "8084"
	fmt.Println("API Gateway running on ", port)
	logger.Info("API Gateway running on " + port)
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error starting server:", err)
	}

	// log.Fatal(http.ListenAndServe(":"+port, nil))
	// if err := http.ListenAndServe(":"+port, mux); err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }
}
