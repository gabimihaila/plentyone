package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	logger "api_gateway/logger"
)

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

func LoadConfig() Config {
	var port string
	var jwtSecret string

	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened config.json")
	logger.Info("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'config' which we defined above
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
