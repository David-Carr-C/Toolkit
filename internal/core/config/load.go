package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func loadConfig(path string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}

func GetProjectDirs(projectName string) ([]string, error) {
	config, err := loadConfig("configs/dir.yaml")
	if err != nil {
		return nil, fmt.Errorf("[GetProjectDirs] error loading config: %w", err)
	}

	if project, exists := config.Projects[projectName]; exists {
		return project.BackupDirs, nil
	}

	return nil, fmt.Errorf("[GetProjectDirs] project %s not found in config", projectName)
}

func GetAllProjects() ([]string, error) {
	config, err := loadConfig("configs/dir.yaml")
	if err != nil {
		return nil, fmt.Errorf("[GetAllProjectNames] error loading config: %w", err)
	}

	var projectNames []string
	for name := range config.Projects {
		projectNames = append(projectNames, name)
	}

	return projectNames, nil
}
