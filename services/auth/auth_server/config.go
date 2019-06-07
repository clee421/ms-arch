package main

import (
	"encoding/json"
	"os"
)

// Configuration type for authentication server configs
type Configuration struct {
	Host     string         `json:"host"`
	Port     int            `json:"port"`
	Database databaseConfig `json:"database"`
	Jwt      jwtConfig      `json:"jwt"`
}

type databaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type jwtConfig struct {
	Issuer     string `json:"iss"`
	Expiration string `json:"exp"`
	Secret     string `json:"secret"`
}

func getServerConfigs(filename string) (*Configuration, error) {
	configuration := Configuration{}

	err := Parse(filename, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

// Parse parses a json path/filename and fills the configuration
func Parse(filename string, configuration *Configuration) error {
	config, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(config)
	err = decoder.Decode(&configuration)
	if err != nil {
		return err
	}

	return nil
}
