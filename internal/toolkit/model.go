package toolkit

import "github.com/spf13/cobra"

func NewToolkit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toolkit",
		Aliases: []string{"tk"},
		Short:   "Aplicación de terminal para la gestión de servidores y su automatización",
		Long:    "Automatiza tareas diarias de administración de servidores, como la creación de backups, la comparación, la actualización y el versionamiento de bases de datos, entre otras.",
		Run:     Run,
	}

	return cmd
}
