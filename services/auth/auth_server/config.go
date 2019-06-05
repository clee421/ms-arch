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

	config, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(config)
	err = decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
