package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenRouterAPIKey string
	Model            string
	ModelFallback    string
}

func New() (*Config, error) {
	cfg := &Config{}
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get working dir, %w", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			break
		}
		parentDir := filepath.Dir(dir)
		if dir == parentDir {
			return nil, fmt.Errorf("find root dir, path: %q", dir)
		}
		dir = parentDir
	}

	envMap, err := godotenv.Read(filepath.Join(dir, ".env"))
	if err != nil {
		return nil, fmt.Errorf("env read, %w", err)
	}

	if port := envMap["OPEN_ROUTER_API_KEY"]; strings.TrimSpace(port) != "" {
		cfg.OpenRouterAPIKey = port
	}

	if env := envMap["MODEL"]; strings.TrimSpace(env) != "" {
		cfg.Model = env
	}

	if env := envMap["MODEL_FALLBACK"]; strings.TrimSpace(env) != "" {
		cfg.ModelFallback = env
	}

	return cfg, nil
}
