package config

import "time"

type InvoiceInterval string

const (
	InvoiceIntervalBimonthly InvoiceInterval = "bimonthly"
	InvoiceIntervalManual    InvoiceInterval = "manual"
)

type TrkConfig struct {
	TimeZone   *time.Location  `yaml:"time_zone"`
	Projects   []ProjectConfig `yaml:"projects"`
	AutoCommit bool            `yaml:"auto_commit"`
}

type ProjectConfig struct {
	Name            string          `yaml:"name"`
	HourlyRate      float64         `yaml:"hourly_rate"`
	InvoiceInterval InvoiceInterval `yaml:"invoice_interval" default:"bimonthly"`
}
