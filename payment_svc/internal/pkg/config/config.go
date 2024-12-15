package config

import (
	"os"
	"sync"
)

type Config struct {
	ServerConfig ServerConfig
	LogFilePath  string
}

type ServerConfig struct {
	Port string
	Host string
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		cfg = &Config{
			ServerConfig: ServerConfig{
				Port: os.Getenv("PAYMENT_HTTP_SERVER_PORT"),
				Host: os.Getenv("PAYMENT_HTTP_SERVER_HOST"),
			},
			LogFilePath: os.Getenv("LOG_FILE_PATH"),
		}
	})

	return cfg, err
}

func LoadServerConfig() (ServerConfig, error) {
	config, err := LoadConfig()
	if err != nil {
		return ServerConfig{}, err
	}

	return config.ServerConfig, nil
}

func LoadLogFilePath() (string, error) {
	config, err := LoadConfig()
	if err != nil {
		return "", err
	}

	return config.LogFilePath, nil
}
