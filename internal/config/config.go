package config

import "time"

type TrkConfig struct {
	TimeZone *time.Location  `yaml:"time_zone"`
	Projects []ProjectConfig `yaml:"projects"`
}

type ProjectConfig struct {
	Name       string  `yaml:"name"`
	HourlyRate float64 `yaml:"hourly_rate"`
}
