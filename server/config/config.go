package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		User      string
		Password  string
		Address   string
		Protocol  string
		Name      string
		ParseTime string
	}
	Server struct {
		Address string
	}
	Environment map[string]string
}

func Read() *Config {
	var c Config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(".", "config"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("failed to parse config file: %w", err))
	}
	for k, v := range c.Environment {
		os.Setenv(strings.ToUpper(k), v)
	}
	spew.Dump(c)
	return &c
}
