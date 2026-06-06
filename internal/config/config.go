package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Formatters []string        `yaml:"formatters"`
	Tasks      map[string]Task `yaml:"tasks"`
}

type Task struct {
	Description string `yaml:"description"`
	Cmd         string `yaml:"cmd"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		path = "dev.yaml"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s not found", path)
		}
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	if cfg.Tasks == nil {
		cfg.Tasks = make(map[string]Task)
	}

	return &cfg, nil
}
