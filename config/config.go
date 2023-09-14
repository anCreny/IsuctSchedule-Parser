package config

import (
	"errors"
	"github.com/anCreny/IsuctSchedule-Packages/structs"
	"os"
)

var (
	NoEnvVarsError = errors.New("no one environment variables were found")
)

var Cfg *Config

type Config struct {
	RxCfg    structs.ReindexerConfig
	ParseUrl string //url to get all schedule to parse
}

func Init() error {
	var cfg = &Config{
		ParseUrl: os.Getenv("URL"),
		RxCfg: structs.ReindexerConfig{
			Host:     os.Getenv("RX_HOST"),
			Port:     os.Getenv("RX_PORT"),
			Username: os.Getenv("RX_USERNAME"),
			Password: os.Getenv("RX_PASSWORD"),
			Database: os.Getenv("RX_DATABASE"),
			Namespaces: structs.Namespaces{
				Teachers: os.Getenv("NM_TEACHERS"),
				Groups:   os.Getenv("NM_GROUPS"),
				Names:    os.Getenv("NM_NAMES"),
			},
		},
	}

	if cSum := cfg.ParseUrl +
		cfg.RxCfg.Host +
		cfg.RxCfg.Port +
		cfg.RxCfg.Password +
		cfg.RxCfg.Username +
		cfg.RxCfg.Database +
		cfg.RxCfg.Namespaces.Teachers +
		cfg.RxCfg.Namespaces.Groups +
		cfg.RxCfg.Namespaces.Names; cSum == "" {
		return NoEnvVarsError
	}

	Cfg = cfg
	return nil
}
