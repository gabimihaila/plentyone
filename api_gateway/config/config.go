package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	logger "api_gateway/logger"
)

var ConfigFile = "config.json"

type Config struct {
	Port         string        `json:"port"`
	JWTSecret    string        `json:"jwt_secret"`
	Destinations []Destination `json:"destinations"`
}

var config Config

type Destination struct {
	URL        string `json:"url"`
	PathPrefix string `json:"path_prefix"`
}

func LoadConfig(configFileName string) Config {
	var port string
	var jwtSecret string

	// Open jsonFile
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened " + configFileName)
	logger.Info("Successfully Opened " + configFileName)
	// close jsonFile at the end of the block
	defer jsonFile.Close()

	// read opened jsonFile as a byte array.
	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	// unmarshal the byteArray which contains 
	// jsonFile's content into 'config'
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}

	port = config.Port
	jwtSecret = config.JWTSecret
	destinationsAll := config.Destinations

	if port == "" {
		log.Fatal("Missing port config.")
	}

	if jwtSecret == "" {
		log.Fatal("Missing jwtSecret config.")
	}

	if len(destinationsAll) == 0 {
		log.Fatal("Missing destinations config.")
	}

	config = Config{
		Port:         port,
		JWTSecret:    jwtSecret,
		Destinations: destinationsAll,
	}

	return config
}
