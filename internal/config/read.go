package config

import (
	"encoding/json"
	"errors"
	"os"
)

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(fullPath)

	if err != nil {
		return Config{}, errors.New("Unable to read from gatorconfig")
	}

	decoder := json.NewDecoder(file)
	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, errors.New("Failed to unmarshal from gatorconfig")
	}

	return cfg, nil
}
