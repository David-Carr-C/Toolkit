package backup

import "github.com/spf13/cobra"

type Flags struct {
	database    string
	source      string
	destination string
}

func NewBackup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup",
		Aliases: []string{"bk"},
		Short:   "Especificación de respaldo",
		Long:    "Automatiza tareas de creación de backups.",
		Run:     Run,
	}

	cmd.Flags().String(
		"database",
		"",
		"Nombre de la base de datos para la que se desea realizar el backup",
	)

	cmd.Flags().String(
		"source",
		"",
		"Origen desde donde se obtendran los datos para el backup",
	)

	cmd.Flags().String(
		"destination",
		"",
		"Ruta de destino donde se guardará el backup",
	)

	return cmd
}
