package backend

import (
	"encoding/json"
	"os"
	"path/filepath"

	"immich-desktop-sync/backend/models"
)

func configDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "immich-desktop")
}

func configPath() string {
	return filepath.Join(configDir(), "config.json")
}

func LoadConfig() (*models.Config, error) {
	data, err := os.ReadFile(configPath())
	if os.IsNotExist(err) {
		return &models.Config{}, nil
	}
	if err != nil {
		return nil, err
	}
	var cfg models.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveConfig(cfg *models.Config) error {
	if err := os.MkdirAll(configDir(), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath(), data, 0600)
}

func DBPath() string {
	return filepath.Join(configDir(), "sync.db")
}
