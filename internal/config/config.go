package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configFileNameBase string = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	//get file path of .gatorconfig.json
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	//open .gatorconfig.json
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	//decode .gatorconfig.json into Config struct and return
	decoder := json.NewDecoder(file)
	cfg := Config{}
	if err = decoder.Decode(&cfg); err != nil {
		return Config{}, err
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


func (cfg *Config) SetUser(username string) error {
	if username == "" {
		return errors.New("username not provided")
	}
	cfg.CurrentUserName = username
	return write(*cfg)
}

func write(cfg Config) error {
	//get file path of .gatorconfig.json
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	//create. os.File typw which satisfies os.writer interface
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	//creates encoder
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}