package toolkit

import "github.com/spf13/cobra"

type Flags struct {
	projectName  string
	databaseName string
}

func NewToolkit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toolkit",
		Aliases: []string{"tk"},
		Short:   "Aplicación de terminal para la gestión de servidores y su automatización",
		Long:    "Automatiza tareas diarias de administración de servidores, como la creación de backups, la comparación, la actualización y el versionamiento de bases de datos, entre otras.",
		Run:     Run,
	}

	cmd.Flags().StringP(
		"project",
		"p",
		"",
		"Nombre del proyecto para el que se desea desglosar información",
	)

	cmd.Flags().StringP(
		"database",
		"d",
		"",
		"Nombre de la base de datos para la que se desea desglosar información",
	)

	return cmd
}
