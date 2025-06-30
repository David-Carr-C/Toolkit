package backup

import "github.com/spf13/cobra"

func NewBackup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup",
		Aliases: []string{"bk"},
		Short:   "Especificación de respaldo",
		Long:    "Automatiza tareas de creación de backups.",
		Run:     Run,
	}

	return cmd
}
