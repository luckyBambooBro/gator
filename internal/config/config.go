package config

import (
	"encoding/json"
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
	cfg.CurrentUserName = username
	//convert cfg into slice of bytes to be written to file
	data, err := write(cfg)
	if err != nil {
		return err
	}
	filepath, err := getConfigFilePath() //obtain filepath of .gatorconfig.json
	if err != nil {
		return err
	}
	os.Create(filepath)
	//write to file (overwrite .gatorconfig.json with new data)
	os.WriteFile(filepath, data, 0644)
	return nil

}

func write(cfg *Config) ([]byte, error){
	
	data, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	return data, nil
}