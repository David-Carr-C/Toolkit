package dir

import (
	"fmt"

	"criteria.mx/scripts/internal/core/config/interfaces"
)

func GetProjectDirs(projectName string) ([]string, error) {
	var config *Config
	err := interfaces.LoadConfig("configs/dir.yaml", &config)
	if err != nil {
		return nil, fmt.Errorf("[GetProjectDirs] error loading config: %w", err)
	}

	if project, exists := config.Projects[projectName]; exists {
		return project.BackupDirs, nil
	}

	return nil, fmt.Errorf("[GetProjectDirs] project %s not found in config", projectName)
}

func GetAllProjects() ([]string, error) {
	var config *Config
	err := interfaces.LoadConfig("configs/dir.yaml", &config)
	if err != nil {
		return nil, fmt.Errorf("[GetAllProjectNames] error loading config: %w", err)
	}

	var projectNames []string
	for name := range config.Projects {
		projectNames = append(projectNames, name)
	}

	return projectNames, nil
}
