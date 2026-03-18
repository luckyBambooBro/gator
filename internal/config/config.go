package config

import (
	"encoding/json"

)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser() error {
	write(cfg)
	fileToWrite, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	//os.WriteFile(string(fileToWrite))
	return nil

}

func write(cfg *Config) {
	cfg.CurrentUserName = "User1" 
	//come back to change this to what the user types in
}