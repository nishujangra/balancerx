package models

type Config struct {
	Port        string   `yaml:"port"`
	Strategy    string   `yaml:"strategy"`
	Backends    []string `yaml:"backends"`
	Protocol    string   `yaml:"protocol"`
	HealthCheck `yaml:"health_check"`
}

type HealthCheck struct {
	Path string `yaml:"path"`
}
