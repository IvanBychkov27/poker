package config

import (
	"fmt"
	"github.com/cristalhq/aconfig"
	"strings"
)

type Config struct {
	MainAddress    string `default:"127.0.0.1:5000" env:"MAIN_ADDRESS"`
	ControlAddress string `default:"127.0.0.1:5001" env:"CONTROL_ADDRESS"`
}

func New() *Config {
	return &Config{}
}

func (cfg *Config) Load() error {
	err := aconfig.LoaderFor(cfg, aconfig.Config{
		SkipFiles: true,
		SkipFlags: true,
		EnvPrefix: "POKER",
	}).Load()
	if err != nil {
		return err
	}

	return cfg.Validate()
}

func (cfg *Config) Validate() error {
	if strings.TrimSpace(cfg.MainAddress) == "" {
		return fmt.Errorf("MainAddress must be defined")
	}

	if strings.TrimSpace(cfg.ControlAddress) == "" {
		return fmt.Errorf("ControlAddress must be defined")
	}

	return nil
}
