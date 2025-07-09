package backup

import (
	"criteria.mx/scripts/internal/core/service/backup"
	"fmt"

	"github.com/spf13/cobra"
)

func getFlags(cmd *cobra.Command) (*Flags, error) {
	database, err := cmd.Flags().GetString("database")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre de la base de datos: %w", err)
	}

	source, err := cmd.Flags().GetString("source")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener la fuente del respaldo: %w", err)
	}

	destination, err := cmd.Flags().GetString("destination")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el destino del respaldo: %w", err)
	}

	return &Flags{
		database,
		source,
		destination,
	}, nil
}

func Run(cmd *cobra.Command, _ []string) {
	flags, err := getFlags(cmd)
	if err != nil {
		fmt.Printf("[Run] Ocurrió el siguiente error: %s\n", err)
		return
	}

	switch {
	case flags.database != "":
		databaseService, err := service.NewDatabaseService(flags.database, flags.source, flags.destination)
		if err != nil {
			fmt.Printf("[Run] Ocurrió el siguiente error al configurar el servicio: %s\n", err)
			return
		}

		err = databaseService.Exec()
		if err != nil {
			fmt.Printf("[Run] Ocurrió el siguiente error al ejecutar el servicio de base de datos: %s\n", err)
			return
		}

		fmt.Println("[Success] Respaldos de la base de datos completados exitosamente.")
	}
}
