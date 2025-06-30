package toolkit

import (
	"fmt"

	"criteria.mx/scripts/internal/core/config"
	"github.com/spf13/cobra"
)

type Flags struct {
	projectName  string
	databaseName string
}

func getFlags(cmd *cobra.Command) (*Flags, error) {
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre del proyecto: %w", err)
	}

	database, err := cmd.Flags().GetString("database")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre de la base de datos: %w", err)
	}

	return &Flags{
		projectName:  project,
		databaseName: database,
	}, nil
}

func seeProjects(projectName string) {
	if projectName == "" {
		fmt.Println("Proyectos:")
		resultSet, _ := config.GetAllProjects()
		for _, project := range resultSet {
			fmt.Print("  - ", project, "\n")
		}
	} else {
		dirs, err := config.GetProjectDirs(projectName)
		if err != nil {
			fmt.Printf("Error al obtener los directorios del proyecto %s: %v\n", projectName, err)
			return
		}

		fmt.Printf("Directorios de respaldo para el proyecto %s:\n", projectName)
		for _, dir := range dirs {
			fmt.Println("  -", dir)
		}
	}
}

func Run(cmd *cobra.Command, args []string) {
	flags, err := getFlags(cmd)
	if err != nil {
		fmt.Printf("Error al obtener las banderas: %v\n", err)
	}

	seeProjects(flags.projectName)
}
