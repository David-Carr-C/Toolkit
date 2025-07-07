package backup

import (
	"criteria.mx/scripts/internal/core/service"
	"fmt"

	"github.com/spf13/cobra"
)

func getFlags(cmd *cobra.Command) (*Flags, error) {
	database, err := cmd.Flags().GetString("database")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre de la base de datos: %w", err)
	}

	return &Flags{
		database,
	}, nil
}

func Run(cmd *cobra.Command, args []string) {
	flags, err := getFlags(cmd)
	if err != nil {
		fmt.Printf("[Run] Ocurrió el siguiente error: %s\n", err)
		return
	}

	switch {
	case flags.database != "":
		err := service.Database(flags.database)
		if err != nil {
			fmt.Printf("[Run] Ocurrió el siguiente error: %s\n", err)
		}
	}
}
