package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileNameBase string = ".gatorconfig.json"

func Read() (*Config, error) {
	fileName, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}

	if err = json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileNameBase), nil
}
