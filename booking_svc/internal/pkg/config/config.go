package config

import (
	"fmt"
	"github.com/spf13/viper"
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
		viper.SetConfigFile("internal/pkg/config/.env")
		viper.AutomaticEnv()

		if e := viper.ReadInConfig(); e != nil {
			err = fmt.Errorf("error reading config file: %v", e)
			return
		}

		cfg = &Config{
			DbConfig: DbConfig{
				Host:     viper.GetString("DB_HOST"),
				Port:     viper.GetString("DB_PORT"),
				Username: viper.GetString("DB_USERNAME"),
				Password: viper.GetString("DB_PASSWORD"),
				Database: viper.GetString("DB_DATABASE"),
			},
			ServerConfig: ServerConfig{
				Port: viper.GetString("SERVER_PORT"),
				Host: viper.GetString("SERVER_HOST"),
			},
			LogFilePath: viper.GetString("LOG_FILE_PATH"),
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
