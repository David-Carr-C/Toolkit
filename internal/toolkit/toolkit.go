package toolkit

import (
	"fmt"

	"criteria.mx/scripts/internal/core/config/db"
	"criteria.mx/scripts/internal/core/config/dir"
	"criteria.mx/scripts/internal/core/config/sv"
	"github.com/spf13/cobra"
)

func getFlags(cmd *cobra.Command) (*Flags, error) {
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre del proyecto: %w", err)
	}

	database, err := cmd.Flags().GetString("database")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre de la base de datos: %w", err)
	}

	server, err := cmd.Flags().GetString("server")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre del servidor: %w", err)
	}

	sync, err := cmd.Flags().GetString("sync")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener la sincronización: %w", err)
	}

	return &Flags{
		projectName:  project,
		databaseName: database,
		serverName:   server,
		sync:         sync,
	}, nil
}

func seeProjects(projectName string) {
	if projectName == "" {
		fmt.Println("Proyectos:")
		resultSet, _ := dir.GetAllProjects()
		for _, project := range resultSet {
			fmt.Print("  - ", project, "\n")
		}
	} else {
		result, err := dir.GetProjectDirs(projectName)
		if err != nil {
			fmt.Printf("Error al obtener los directorios del proyecto %s: %v\n", projectName, err)
			return
		}

		fmt.Printf("Directorios de respaldo para el proyecto %s:\n", projectName)
		for _, dir := range result {
			fmt.Println("  -", dir)
		}
	}
}

func seeDatabases(databaseName string) {
	if databaseName == "" {
		fmt.Println("Bases de datos:")
		resultSet, _ := db.GetAllDatabases()
		for _, db := range resultSet {
			fmt.Printf("  - %s: %s\n", db["database"], db["status"])
		}
	} else {
		dirs, err := db.GetDatabaseConfig(databaseName)
		if err != nil {
			fmt.Printf("Error al obtener las bases de datos %s: %v\n", databaseName, err)
			return
		}

		fmt.Printf("Bases de datos para %s:\n", databaseName)
		for _, dir := range dirs {
			fmt.Println("  -", dir)
		}
	}

}

func seeServers(serverName string) {
	if serverName == "" {
		fmt.Println("Servidores:")
		resultSet, _ := sv.GetAllServers()
		for _, server := range resultSet {
			fmt.Printf("  - %s: %s\n", server["server"], server["configured"])
		}
	} else {
		host, path, err := sv.GetServerPath(serverName)
		if err != nil {
			fmt.Printf("Error al obtener el servidor %s: %v\n", serverName, err)
			return
		}

		fmt.Printf("Servidor %s:\n", host)
		fmt.Println("  - Path:", path)
	}
}

func sync(server string) {
	// TODO: Implementación hacia Elasticsearch
}

func Run(cmd *cobra.Command, args []string) {
	flags, err := getFlags(cmd)
	if err != nil {
		fmt.Printf("Error al obtener las banderas: %v\n", err)
	}

	if flags.projectName == "" && flags.databaseName == "" && flags.sync == "" && flags.serverName == "" {
		fmt.Println("Visualiza un proyecto o base de datos específica con \"-h\"")
		seeProjects(flags.projectName)
		seeDatabases(flags.databaseName)
		seeServers(flags.serverName)
		return
	}

	switch {
	case flags.sync != "":
		sync(flags.sync)
	case flags.projectName != "":
		seeProjects(flags.projectName)
	case flags.databaseName != "":
		seeDatabases(flags.databaseName)
	case flags.serverName != "":
		seeServers(flags.serverName)
	}
}
