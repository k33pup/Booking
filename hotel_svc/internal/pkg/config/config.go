package config

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"slices"
)

const (
	defaultMaxHeaderBytes = 1 << 20

	ParamGRCPAddress     = "GRPC_ADDRESS"
	ParamHTTPAddress     = "HTTP_ADDRESS"
	ParamNameHTTPTimeout = "HTTP_TIMEOUT"
	ParamLogPath         = "LOG_PATH"
	ParamGRCPPort        = "GRPC_SERVER_PORT"
	ParamHTTPPort        = "HTTP_SERVER_PORT"
	ParamDSNString       = "DSB_STRING"
)

var (
	ErrParamNotFound = errors.New("param not found")
	ErrBadConfig     = errors.New("bad config")
)

type Config struct {
	params map[string]string
}

func NewConfig() *Config {
	return &Config{
		params: map[string]string{
			ParamGRCPAddress:     ":" + os.Getenv(ParamGRCPPort),
			ParamHTTPAddress:     ":" + os.Getenv(ParamHTTPPort),
			ParamNameHTTPTimeout: "10s",
			ParamLogPath:         "/tmp/log.txt",
			ParamDSNString: fmt.Sprintf(
				"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				os.Getenv("POSTGRES_HOST"),
				os.Getenv("POSTGRES_PORT"),
				os.Getenv("POSTGRES_USER"),
				os.Getenv("POSTGRES_DB"),
				os.Getenv("POSTGRES_PASSWORD"),
			),
		},
	}
}

func (c *Config) setParam(key, value string) error {
	if _, ok := c.params[key]; !ok {
		return fmt.Errorf("%w: %s", ErrParamNotFound, key)
	}
	c.params[key] = value
	return nil
}

func (c *Config) getParamNames() []string {
	return slices.Collect(maps.Keys(c.params))
}

func (c *Config) GetGRCPAddress() string {
	return c.params[ParamGRCPAddress]
}

func (c *Config) GetHTTPAddress() string {
	return c.params[ParamHTTPAddress]
}

func (c *Config) GetTimeout() string {
	return c.params[ParamNameHTTPTimeout]
}

func (c *Config) GetLogPath() string {
	return c.params[ParamLogPath]
}

func (c *Config) GetDSN() string {
	return c.params[ParamDSNString]
}
