package request_manager

import (
	"testing"
	"time"
)

func TestLogRequest(t *testing.T) {
	expectedEndpoint := "http://192.168.1.999:8086/serv9"
	expectedMethod := "GET"

	LogRequest(expectedMethod, expectedEndpoint)
}

func TestLogLatency(t *testing.T) {
	expectedEndpoint := "http://192.168.1.999:8086/serv9"
	start := time.Now()
	expectedDuration := time.Since(start).Seconds()

	LogLatency(expectedEndpoint, time.Duration(expectedDuration))
}
