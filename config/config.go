package config

import (
	"errors"
	"os"
)

var (
	NoEnvVarsError = errors.New("no one environment variables were found")
)

var Cfg *Config

type Config struct {
	RxCfg    *RxConfig
	ParseUrl string //url to get all schedule to parse
}

// config to get connection to reindexer database
type RxConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func Init() error {
	var cfg = &Config{
		ParseUrl: os.Getenv("URL"),
		RxCfg: &RxConfig{
			Host:     os.Getenv("RX_HOST"),
			Port:     os.Getenv("RX_PORT"),
			Username: os.Getenv("RX_USERNAME"),
			Password: os.Getenv("RX_PASSWORD"),
			Database: os.Getenv("RX_DATABASE"),
		},
	}

	cSum := cfg.ParseUrl + cfg.RxCfg.Host + cfg.RxCfg.Port + cfg.RxCfg.Password + cfg.RxCfg.Username + cfg.RxCfg.Database
	if cSum == "" {
		return NoEnvVarsError
	}

	Cfg = cfg
	return nil
}
