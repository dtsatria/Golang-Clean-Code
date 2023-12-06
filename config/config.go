package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type TokenConfig struct {
	IssuerName      string
	JwtSignatureKey []byte
	JwtLifeTime     time.Duration
}
type LogFileConfig struct {
	FilePath string
}

type EmailConfig struct {
	Server    string
	Port      string
	EmailFrom string
	Password  string
}

type Config struct {
	ApiConfig
	EmailConfig
	DbConfig
	TokenConfig
	LogFileConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.EmailConfig = EmailConfig{
		Server:    os.Getenv("EMAIL_SERVER"),
		Port:      os.Getenv("EMAIL_PORT"),
		EmailFrom: os.Getenv("EMAIL_FROM"),
		Password:  os.Getenv("EMAIL_PASSWORD"),
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.LogFileConfig = LogFileConfig{
		FilePath: os.Getenv("LOG_FILE"),
	}

	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		return err
	}

	c.TokenConfig = TokenConfig{
		IssuerName:      os.Getenv("TOKEN_ISSUE_NAME"),
		JwtSignatureKey: []byte(os.Getenv("TOKEN_KEY")),
		JwtLifeTime:     time.Duration(tokenLifeTime) * time.Hour,
	}

	if c.ApiPort == "" || c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" ||
		c.DbConfig.Password == "" || c.Driver == "" || c.EmailConfig.Server == "" || c.EmailConfig.Port == "" || c.EmailConfig.EmailFrom == "" ||
		c.EmailConfig.Password == "" {
		return errors.New("missing required environment variables")
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
