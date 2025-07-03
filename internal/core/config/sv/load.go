package sv

import (
	"fmt"
	"os/exec"

	"criteria.mx/scripts/internal/core/config/interfaces"
)

func GetServerPath(projectName string) (string, string, error) {
	var config *Config
	err := interfaces.LoadConfig("configs/sv.yaml", &config)
	if err != nil {
		return "", "", fmt.Errorf("[GetProjectDirs] error loading config: %w", err)
	}

	if server, ok := config.Servers[projectName]; ok {
		if server.Path != "" {
			return server.Host, server.Path, nil
		}
		return "", "", fmt.Errorf("[GetProjectDirs] project %s has no path configured", projectName)
	}

	return "", "", fmt.Errorf("[GetProjectDirs] project %s not found in config", projectName)
}

func testConnection(host string) error {
	if host == "localhost" {
		return nil
	}

	cmd := exec.Command("ssh", "-o", "BatchMode=yes", host, "exit")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[testConnection] error connecting to host %s: %w", host, err)
	}

	return nil
}

func GetAllServers() ([]map[string]string, error) {
	var config *Config
	err := interfaces.LoadConfig("configs/sv.yaml", &config)
	if err != nil {
		return nil, fmt.Errorf("[GetAllServers] error loading config: %w", err)
	}

	var servers []map[string]string
	for name, server := range config.Servers {
		err := testConnection(server.Host)
		configured := "configured"
		if err != nil {
			configured = "error"
		}

		servers = append(servers, map[string]string{
			"server":     name,
			"configured": configured,
		})
	}

	return servers, nil
}
