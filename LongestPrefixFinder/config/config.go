package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

const CONFIG_FILE = "config/config.json"
const CONFIG_FILE_ERROR = "ERROR: config.json file not present in current directory"

var appConfig config

// Loading json config file.
func LoadAppConfig() {

	if _, err := os.Stat(CONFIG_FILE); errors.Is(err, os.ErrNotExist) {
		log.Fatal(CONFIG_FILE_ERROR)
	}

	jsonData, _ := os.ReadFile(CONFIG_FILE)
	err := json.Unmarshal(jsonData, &appConfig)
	if err != nil {
		log.Fatal(err)
	}
}

// Function to get prefix file path value.
func GetPrefixFilePath() string {
	return appConfig.PrefixFilePath
}

// Function to get input value.
func GetInputValue() string {
	return appConfig.InputValue
}

// Struct to hold application configurations.
type config struct {
	PrefixFilePath string `json:"prefix_file_path"`
	InputValue     string `json:"input_value"`
}
