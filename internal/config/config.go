package config

type TrkConfig struct {
	Projects []ProjectConfig `yaml:"projects"`
}

type ProjectConfig struct {
	Name       string  `yaml:"name"`
	HourlyRate float32 `yaml:"hourly_rate"`
}
