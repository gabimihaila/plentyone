package handler

import (
	"api_gateway/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var expectedConfigFile = "../config-test.json"

func TestGetTargetUrl(t *testing.T) {
	expectedDestinations := []config.Destination{
		{
			URL:        "http://192.168.1.999:8086",
			PathPrefix: "serv25",
		},
	}

	expectedPath := "serv25"

	expectedUrl := "http://192.168.1.999:8086"

	url := getTargetUrl(expectedDestinations, expectedPath)

	if url != expectedUrl {
		t.Errorf("Expected url %s, got %s", expectedUrl, url)
	}
}

var ExpectedConfig = config.LoadConfig(expectedConfigFile)

var ExpectedAPIBaseURL = getEnv("API_BASE_URL", "http://192.168.1.999:8086/serv25/home")

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func TestProxyHandlerStatusBadRequest(t *testing.T) {
	expectedDestinations := []config.Destination{
		{
			URL:        "http://192.168.1.999:8086",
			PathPrefix: "serv25",
		},
		{
			URL:        "http://192.168.1.999:8086",
			PathPrefix: "serv15",
		},
	}

	if len(expectedDestinations) == 0 {
		t.Errorf("Expected destinations is empty")
	}

	handler := ProxyHandler(expectedDestinations)

	test := struct {
		name         string
		path         string
		expectStatus int
	}{
		"Invalid Path", "/invalid", http.StatusBadRequest,
	}

	t.Run(test.name, func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, test.path, nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)

		if recorder.Code != test.expectStatus {
			t.Errorf("Expected status %d, got %d", test.expectStatus, recorder.Code)
		}
	})
}
