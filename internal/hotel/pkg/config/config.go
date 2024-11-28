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
			ParamGRCPAddress:     "localhost:9090",
			ParamHTTPAddress:     "localhost:9091",
			ParamNameHTTPTimeout: "10s",
			ParamLogPath:         "/tmp/log.txt",
		},
	}
}

func (c *Config) InitFromEnv() error {
	for _, k := range c.getParamNames() {
		v := os.Getenv(k)
		if v != "" {
			err := c.setParam(k, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
