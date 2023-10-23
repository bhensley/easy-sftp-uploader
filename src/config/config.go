package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Hosts []struct {
		HostAddress string `json:"addr"`
		Post        string `json:"port"`
		Username    string `json:"username"`
	} `json:"hosts"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)

	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
