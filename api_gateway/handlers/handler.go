package handler

import (
	"api_gateway/config"
	logger "api_gateway/logger"
	"api_gateway/request_manager"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func getTargetUrl(destinationsAll []config.Destination, path string) string {
	var url string
	var path_prefix string

	for i := 0; i < len(destinationsAll); i++ {
		destination := destinationsAll[i]
		fmt.Println("destination  ", destination)
		url = destination.URL
		path_prefix = destination.PathPrefix

		if path_prefix == path {
			return url
		}
	}

	return ""
}

// ProxyHandler forwards requests to a target service
func ProxyHandler(destinationsConfig []config.Destination) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		destinationsConfig := config.LoadConfig(config.ConfigFile).Destinations
		path := strings.Split(r.URL.Path, "/")[1]

		url := getTargetUrl(destinationsConfig, path)

		if url == "" {
			http.Error(w, "Destination not found", http.StatusBadRequest)
			logger.Error(errors.New("Destination not found for path: " + path))
			return
		}

		targetURL := url
		fmt.Println("targetURL ", targetURL)

		resp, err := http.Get(targetURL)
		if err != nil {
			http.Error(w, "Error forwarding request", http.StatusBadGateway)
			logger.Error(err)
			return
		}
		defer resp.Body.Close()

		duration := time.Since(start).Seconds()

		fullURL := targetURL + r.URL.Path

		request_manager.LogRequest(r.Method, targetURL)
		request_manager.LogLatency(fullURL, time.Duration(duration))

		fmt.Println("Request to ", fullURL, " took ", duration, " seconds ")
		infoMessage := fmt.Sprintf("Request to %s took %.7f seconds", fullURL, duration)
		logger.Info(infoMessage)

		w.WriteHeader(resp.StatusCode)

		infoMessage = fmt.Sprintf("[API GATEWAY] Method: %s | Path: %s | Status: %s | Duration: %v",
			r.Method, fullURL, resp.Status, duration)

		logger.Info(infoMessage)

		fmt.Println("resp status code:  ", resp.StatusCode)
		fmt.Println("resp status Name?:  ", resp.Status)
		fmt.Println("r method:  ", r.Method)    

		fmt.Fprint(w, "Request relayed successfully")
		fmt.Println("Request relayed successfully")
		logger.Info("Request relayed successfully")
	}
}
