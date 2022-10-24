package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		User      string
		Password  string
		Address   string
		Port      string
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
		log.Println(fmt.Errorf("failed to read config file: %w", err))
		return nil
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Println(fmt.Errorf("failed to parse config file: %w", err))
		return nil
	}
	for k, v := range c.Environment {
		os.Setenv(strings.ToUpper(k), v)
	}
	return &c
}
