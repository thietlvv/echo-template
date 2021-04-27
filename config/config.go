package config

import (
	"os"
	"strconv"
)

type ConfigSource interface {
	GetString(name string) string
}

type Config struct {
	ConfigSource
}

type EnvGetter struct{}

func (c *Config) GetString(name string) string {
	if nil == c.ConfigSource {
		return ""
	}
	return c.ConfigSource.GetString(name)
}
func (c *Config) GetBool(name string) bool {
	s := c.GetString(name)
	i, err := strconv.ParseBool(s)
	if nil != err {
		return false
	}
	return i
}
func (c *Config) GetInt(name string) int64 {
	s := c.GetString(name)
	i, err := strconv.ParseInt(s, 10, 0)
	if nil != err {
		return 0
	}
	return i
}
func (c *Config) GetFloat(name string) float64 {
	s := c.GetString(name)
	i, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0
	}
	return i
}

func (r *EnvGetter) GetString(name string) string {
	return os.Getenv(name)
}
