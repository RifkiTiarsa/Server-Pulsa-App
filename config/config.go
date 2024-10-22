package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}

type TokenConfig struct {
	IssuerName       string `json:"IssuerName"`
	JwtSignatureKy   []byte `json:"JwtSignatureKy"`
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtExpiresTime   time.Duration
}

type Config struct {
	DbConfig
	ApiConfig
	TokenConfig
}

func (c *Config) readConfigEnvironment() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing .env file %v", err.Error())
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{ApiPort: os.Getenv("API_PORT")}

	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	c.TokenConfig = TokenConfig{
		IssuerName:       os.Getenv("TOKEN_ISSUE"),
		JwtSignatureKy:   []byte(os.Getenv("TOKEN_SECRET")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtExpiresTime:   time.Duration(tokenExpire) * time.Minute,
	}

	if c.Host == "" || c.Port == "" || c.User == "" || c.Name == "" || c.Driver == "" || c.ApiPort == "" ||
		c.IssuerName == "" || c.JwtExpiresTime < 0 || len(c.JwtSignatureKy) == 0 {
		return fmt.Errorf("missing required environment")
	}

	return nil

}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfigEnvironment(); err != nil {
		return nil, err
	}

	return cfg, nil
}
