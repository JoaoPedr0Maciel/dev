package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tasks map[string]Task `yaml:"tasks"`
}

type Task struct {
	Description string `yaml:"description"`
	Cmd         string `yaml:"cmd"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("dev.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("dev.yaml not found")
		}
		return nil, fmt.Errorf("reading dev.yaml: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing dev.yaml: %w", err)
	}

	if cfg.Tasks == nil {
		cfg.Tasks = make(map[string]Task)
	}

	return &cfg, nil
}
