package db

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func expandEnvVars(s string) string {
	re := regexp.MustCompile(`\${(\w+)}`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		varName := strings.TrimPrefix(match, "${")
		varName = strings.TrimSuffix(varName, "}")
		if value, exists := os.LookupEnv(varName); exists {
			return value
		}
		return match
	})
}

func isEnvVarRaw(s string) bool {
	re := regexp.MustCompile(`^\${\w+}$`)
	return re.MatchString(s)
}

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

func GetAllDatabases() ([]map[string]string, error) {
	config, err := loadConfig("configs/db.yaml")
	if err != nil {
		return nil, fmt.Errorf("[GetAllDatabases] error loading config: %w", err)
	}

	var dbStatus []map[string]string
	for key, db := range config.Databases {
		db.Password = expandEnvVars(db.Password)
		var configured string

		if isEnvVarRaw(db.Password) {
			configured = "not configured"
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
