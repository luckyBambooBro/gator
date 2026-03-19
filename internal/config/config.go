package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
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