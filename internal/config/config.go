package config

import (
	"errors"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Unable to retrieve User Home Directory")
	}
	return (home + "/.gatorconfig.json"), nil
}

func (c *Config) SetUser(username string) error {

	c.CurrentUsername = username
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}
