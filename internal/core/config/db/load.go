package db

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func loadConfig(path string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("[loadConfig] Error opening config file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("[loadConfig] Error decoding config file: %w", err)
	}
	return config, nil
}

func GetDatabaseConfig(dbName string) ([]string, error) {
	config, err := loadConfig("configs/db.yaml")
	if err != nil {
		return nil, fmt.Errorf("[GetDatabaseConfig] error loading config: %w", err)
	}

	if dbConfig, exists := config.Databases[dbName]; exists {
		return dbConfig.BackupDbs, nil
	}

	return nil, fmt.Errorf("[GetDatabaseConfig] database %s not found in config", dbName)
}

func GetAllDatabases() ([]string, error) {
	config, err := loadConfig("configs/db.yaml")
	if err != nil {
		return nil, fmt.Errorf("[GetAllDatabases] error loading config: %w", err)
	}

	var dbNames []string
	for name := range config.Databases {
		dbNames = append(dbNames, name)
	}

	return dbNames, nil
}
