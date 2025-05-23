package models

type Config struct {
	Port        string   `yaml:"port"`
	Strategy    string   `yaml:"strategy"`
	Backends    []string `yaml:"backends"`
	Protocol    string   `yaml:"protocol"`
	HealthCheck struct {
		Interval string `yaml:"interval"`
		Path     string `yaml:"path"`
	} `yaml:"health_check"`
}
