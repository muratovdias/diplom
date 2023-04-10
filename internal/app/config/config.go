package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Name string
	Pwd  string
	User string
	Port string
	Host string
}

// GetConfig parse .env file and fill GlobalConfig struct
func GetConfig() (*DBConfig, error) {
	if err := godotenv.Load(); err != nil {
		return &DBConfig{}, err
	}
	config := &DBConfig{
		Name: os.Getenv("DB_NAME"),
		Pwd:  os.Getenv("DB_PWD"),
		User: os.Getenv("DB_USER"),
		Port: os.Getenv("DB_PORT"),
		Host: os.Getenv("DB_HOST"),
	}
	return config, nil
}
