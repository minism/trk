package storage

import (
	"os"

	"github.com/minism/trk/internal/config"
	"gopkg.in/yaml.v3"
)

func SaveConfig(cfg config.TrkConfig) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(config.GetConfigPath(), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() (config.TrkConfig, error) {
	data, err := os.ReadFile(config.GetConfigPath())
	if err != nil {
		return config.TrkConfig{}, err
	}
	var cfg config.TrkConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config.TrkConfig{}, err
	}
	return cfg, nil
}
