package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration for the API Gateway.
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Routes []struct {
		Path         string `yaml:"path"` // e.g., "/api/v1"
		Target       string `yaml:"target"`
		RefillRate   int    `yaml:"refillRate"`
		Capacity     int    `yaml:"capacity"`
		AuthRequired bool   `yaml:"authRequired"`
	} `yaml:"routes"`
	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

// LoadConfig reads a YAML config file from the given path, expands environment variables,
// unmarshals it into a Config struct, and applies default values where needed.
func LoadConfig() (*Config, error) {
	// Read the file using os.ReadFile.
	data, err := os.ReadFile("config/configs.yaml")
	if err != nil {
		return nil, err
	}

	// Expand environment variables in the config file content.
	expandedData := os.ExpandEnv(string(data))

	// Unmarshal the YAML data into our Config struct.
	var cfg Config
	if err := yaml.Unmarshal([]byte(expandedData), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
