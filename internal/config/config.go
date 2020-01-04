package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	TLS      bool   `json:"tls"`
	Cert     string `json:"cert"`
	Key      string `json:"key"`
	HTTP2    bool   `json:"http_2"`
	Port     int    `json:"port"`
	Address  string `json:"address"`
	Database string `json:"database"`
}

func New() (Config, error) {
	configPath := "config.json" //hardcode path if command line param not set

	contents, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("loading config at %s: %w", configPath, err)
	}

	var config Config

	if err := json.Unmarshal(contents, &config); err != nil {
		return Config{}, fmt.Errorf("loading config at %s: %w", configPath, err)
	}

	return config, nil
}
