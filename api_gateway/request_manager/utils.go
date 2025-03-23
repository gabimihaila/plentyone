package request_manager

import (
	"fmt"
	"log"
	"sync"
	"time"
	"api_gateway/logger"
)

// RequestMetrics stores request counts and latencies
type RequestMetrics struct {
	mu            sync.Mutex
	requestCounts map[string]int
	latencies     map[string][]time.Duration
}

// Global metrics instance
var Metrics = RequestMetrics{
	requestCounts: make(map[string]int),
	latencies:     make(map[string][]time.Duration),
}

// LogRequest tracks the number of requests per endpoint
func LogRequest(method, endpoint string) {
	Metrics.mu.Lock()
	defer Metrics.mu.Unlock()
	key := method + " " + endpoint
	Metrics.requestCounts[key]++
	log.Printf("[Metrics] Request: %s | Count: %d", key, Metrics.requestCounts[key])
	infoMessage := fmt.Sprintf("[Metrics] Request: %s | Count: %d", key, Metrics.requestCounts[key])
	logger.Info(infoMessage)
}

// LogLatency tracks the request processing time per endpoint
func LogLatency(endpoint string, duration time.Duration) {
	Metrics.mu.Lock()
	defer Metrics.mu.Unlock()
	Metrics.latencies[endpoint] = append(Metrics.latencies[endpoint], duration)
	log.Printf("[Metrics] Latency: %s | Duration: %v", endpoint, duration)
	infoMessage := fmt.Sprintf("[Metrics] Latency: %s | Duration: %v", endpoint, duration)
	logger.Info(infoMessage)
}
