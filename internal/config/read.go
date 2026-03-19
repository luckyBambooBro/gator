package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (*Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	//read .gatorconfig.json
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	//unmarshal and return .gatorconfig.json
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
