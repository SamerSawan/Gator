package config

import (
	"encoding/json"
	"errors"
	"os"
)

func write(cfg Config) error {
	fullPath, _ := getConfigFilePath()

	file, err := os.Create(fullPath)
	if err != nil {
		return errors.New("Failed to open config file during write")
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return errors.New("Failed to encode json data")
	}

	return nil
}
