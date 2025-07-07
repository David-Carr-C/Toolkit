package db

import (
	"fmt"

	"criteria.mx/scripts/internal/core/config/interfaces"
	"criteria.mx/scripts/pkg"
)

func GetDatabaseStruct(dbName string) (*DatabaseConfig, error) {
	var config *Config
	err := interfaces.LoadConfig("configs/db.yaml", &config)
	if err != nil {
		return nil, fmt.Errorf("[GetDatabaseConfig] error loading config: %w", err)
	}

	if dbConfig, exists := config.Databases[dbName]; exists {
		dbConfig.Password = pkg.ExpandEnvVars(dbConfig.Password)
		return &dbConfig, nil
	}

	return nil, fmt.Errorf("[GetDatabaseConfig] database %s not found in config", dbName)
}

func GetDatabaseConfig(dbName string) ([]string, error) {
	var config *Config
	err := interfaces.LoadConfig("configs/db.yaml", &config)
	if err != nil {
		return nil, fmt.Errorf("[GetDatabaseConfig] error loading config: %w", err)
	}

	if dbConfig, exists := config.Databases[dbName]; exists {
		return dbConfig.BackupDbs, nil
	}

	return nil, fmt.Errorf("[GetDatabaseConfig] database %s not found in config", dbName)
}

func GetAllDatabases() ([]map[string]string, error) {
	var config *Config
	err := interfaces.LoadConfig("configs/db.yaml", &config)
	if err != nil {
		return nil, fmt.Errorf("[GetAllDatabases] error loading config: %w", err)
	}

	var dbStatus []map[string]string
	for key, db := range config.Databases {
		db.Password = pkg.ExpandEnvVars(db.Password)
		var configured string

		if pkg.IsEnvVarRaw(db.Password) {
			configured = "error"
		} else {
			configured = "configured"
		}

		dbStatus = append(dbStatus, map[string]string{
			"database": key,
			"status":   configured,
		})
	}

	return dbStatus, nil
}
