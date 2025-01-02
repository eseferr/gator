package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Unmarshal JSON config file
func Read()(Config, error){
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	config, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}
	var data Config
	err = json.Unmarshal(config, &data)
	if err != nil {
	   return Config{}, err
	}
   return data, nil
   }
   
   func (c *Config) SetUser(currentUser string)error{
	c.CurrentUserName = currentUser
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
   }

   //Get config filepath from home directory
   func getConfigFilePath() (string, error){
	homeDir, err := os.UserHomeDir() 
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil
   }
   func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
    if err != nil {
        return err
    }
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(fullPath,jsonData,0600)
	if err != nil {
		return err
	}	
	return nil
   }