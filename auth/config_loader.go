package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	RecommendedConfigPathFlag   = "config"
	RecommendedConfigPathEnvVar = "SECRETCTL_CONFIG"
	RecommendedHomeDir          = ".secretctl"
	RecommendedFileName         = "config.yaml"
)

// 1. explicit flags
// 2. env var
// 3. ~/.secretctl/config.yaml
func ResolveConfigPath(explicitPath string) (string, error) {

	if explicitPath != "" {
		return explicitPath, nil
	}

	if p := os.Getenv(RecommendedConfigPathEnvVar); p != "" {
		return p, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Cannot determine hom dir: %w", err)
	}

	return filepath.Join(home, RecommendedHomeDir, RecommendedFileName), nil
}

func LoadConfig(explicitPath string) (*Config, error) {

	path, err := ResolveConfigPath(explicitPath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)

	if err != nil {
		if os.IsExist(err) {
			return nil, fmt.Errorf("config file not found: %s\nRun 'secretctl init' to create one", path)
		}
		fmt.Errorf("Cannot read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config format: %w", err)
	}

	return &cfg, nil
}
