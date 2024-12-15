package config

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	DbConfig     DbConfig
	ServerConfig ServerConfig
	LogFilePath  string
}

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
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
			DbConfig: DbConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     os.Getenv("POSTGRES_PORT"),
				Username: os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				Database: os.Getenv("POSTGRES_DB"),
			},
			ServerConfig: ServerConfig{
				Port: os.Getenv("SERVER_PORT"),
				Host: os.Getenv("SERVER_HOST"),
			},
			LogFilePath: os.Getenv("LOG_FILE_PATH"),
		}
	})

	return cfg, err
}

func Dsn() (string, error) {
	config, err := LoadConfig()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.Username,
		config.DbConfig.Database,
		config.DbConfig.Password,
	), nil
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
