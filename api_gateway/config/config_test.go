package config

import (
	"os"
	"testing"
)

// func createTempConfigFile(t *testing.T, content []byte) string {
// 	t.Helper()
// 	file, err := os.Create("config-test.json")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp file: %v", err)
// 	}

// 	_, err = file.Write(content)
// 	if err != nil {
// 		t.Fatalf("Failed to write to temp file: %v", err)
// 	}

// 	file.Close()
// 	return file.Name()
// }

var expectedConfigFile = "../config-test.json"

func TestLoadConfig(t *testing.T) {
	expectedDestination := Destination{
		URL:        "http://192.168.1.999:8086",
		PathPrefix: "serv25",
	}

	expectedConfig := Config{
		Port:      "8080",
		JWTSecret: "mysecretverysecretsupersecretmaxi",
		Destinations: []Destination{
			{URL: "http://192.168.1.999:8086", PathPrefix: "serv15"},
			expectedDestination,
		},
	}

	jsonFile, err := os.Open(expectedConfigFile)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	loadedConfig := LoadConfig(expectedConfigFile)

	if loadedConfig.Port != expectedConfig.Port {
		t.Errorf("Expected port %s, got %s", expectedConfig.Port, loadedConfig.Port)
	}

	if loadedConfig.JWTSecret != expectedConfig.JWTSecret {
		t.Errorf("Expected JWTSecret %s, got %s", expectedConfig.JWTSecret, loadedConfig.JWTSecret)
	}

	if len(loadedConfig.Destinations) != len(expectedConfig.Destinations) {
		t.Fatalf("Expected %d destinations, got %d", len(expectedConfig.Destinations), len(loadedConfig.Destinations))
	}

	for i, dest := range loadedConfig.Destinations {
		if dest.URL != expectedConfig.Destinations[i].URL || dest.PathPrefix != expectedConfig.Destinations[i].PathPrefix {
			t.Errorf("Expected destination %+v, got %+v", expectedConfig.Destinations[i], dest)
		}
	}
}
