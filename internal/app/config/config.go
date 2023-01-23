package config

import "github.com/kamalshkeir/kenv"

type GlobalConfig struct {
	Port  string `kenv:"PORT|:8888"`
	PsgDB struct {
		Name string `kenv:"DB_NAME|db"` // NOT REQUIRED: if DB_NAME not found, defaulted to 'db'
		Pwd  string `kenv:"DB_PWD"`     // REEQUIRED: this env var is required, you will have error if empty
		User string `kenv:"DB_USR"`
		Port string `kenv:"DB_PORT"` // NOT REQUIRED: if DB_DSN not found it's not required, it's ok to stay empty
	}
}

// GetConfig parse .env file and fill GlobalConfig struct
func GetConfig() (*GlobalConfig, error) {
	kenv.Load(".env")
	Config := &GlobalConfig{}
	if err := kenv.Fill(Config); err != nil {
		return Config, err
	}
	return Config, nil
}
