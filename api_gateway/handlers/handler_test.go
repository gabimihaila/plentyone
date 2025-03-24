package handler

import (
	"api_gateway/config"
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
