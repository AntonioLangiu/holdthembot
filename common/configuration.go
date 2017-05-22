package common

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	TelegramAPI    string
    TempFolder     string
}

func LoadConfiguration() *Configuration {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Printf("error, invalid config file")
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Printf("error:", err)
	}
	// Get the variables also from the environment
	token := os.Getenv("TelegramAPI")
	if token != "" {
		configuration.TelegramAPI = token
	}
	token = os.Getenv("TempFolder")
	if token != "" {
		configuration.TempFolder = token
	}
	return &configuration
}
