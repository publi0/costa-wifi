package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".costa-wifi"

type Config struct {
	Data map[string]string `json:"data"`
}

func WriteConfigValue(key, value string) error {
	config, err := readConfigFile()
	if err != nil {
		return fmt.Errorf("error reading existing config: %w", err)
	}

	config.Data[key] = value

	return writeConfigFile(config)
}

func ReadConfigValue(key string) (string, error) {
	config, err := readConfigFile()
	if err != nil {
		return "", fmt.Errorf("error reading config: %w", err)
	}

	value, exists := config.Data[key]
	if !exists {
		return "", fmt.Errorf("key '%s' not found in config", key)
	}

	return value, nil
}

func readConfigFile() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("error getting home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, configFileName)

	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{Data: make(map[string]string)}, nil
		}
		return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("error decoding config data: %w", err)
	}

	return config, nil
}

func writeConfigFile(config Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, configFileName)

	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") 
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("error encoding config data: %w", err)
	}

	return nil
}
