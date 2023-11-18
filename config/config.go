package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	Driver   string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	if c.ApiPort == "" || c.Host == "" || c.Port == "" || c.User == "" || c.Password == "" || c.DbName == "" {
		return errors.New("environment required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
