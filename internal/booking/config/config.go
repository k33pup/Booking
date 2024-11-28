package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadDbConfig() (DbConfig, error) {
	viper.SetConfigFile("internal/booking/config/.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return DbConfig{}, err
	}

	return DbConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		Database: viper.GetString("DB_DATABASE"),
	}, nil
}

func GetDsnString() (string, error) {
	dbConfig, err := LoadDbConfig()
	if err != nil {
		return "", err
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Database,
		dbConfig.Password,
	)

	return dsn, nil
}

type ServerConfig struct {
	Port string
	Host string
}

func LoadServerConfig() (ServerConfig, error) {
	viper.SetConfigFile("internal/booking/config/.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return ServerConfig{}, err
	}

	return ServerConfig{
		Port: viper.GetString("SERVER_PORT"),
		Host: viper.GetString("SERVER_HOST"),
	}, nil
}
