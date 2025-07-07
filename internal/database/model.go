package database

import "github.com/spf13/cobra"

type Flags struct {
	database string
}

func NewDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "database",
		Aliases: []string{"tk"},
		Short:   "Especificación de versionamiento, comparación y actualización de bases de datos",
		Long:    "Automatiza tareas de administración de bases de datos.",
		Run:     Run,
	}

	cmd.Flags().StringP(
		"database",
		"d",
		"",
		"Base de datos a administrar",
	)

	return cmd
}
