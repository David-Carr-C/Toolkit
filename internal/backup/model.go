package backup

import "github.com/spf13/cobra"

type Flags struct {
	database string
}

func NewBackup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup",
		Aliases: []string{"bk"},
		Short:   "Especificación de respaldo",
		Long:    "Automatiza tareas de creación de backups.",
		Run:     Run,
	}

	cmd.Flags().StringP(
		"database",
		"d",
		"",
		"Nombre de la base de datos para la que se desea realizar el backup",
	)

	return cmd
}
