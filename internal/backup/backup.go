package backup

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Flags struct {
	databaseName string
}

func getFlags(cmd *cobra.Command) (*Flags, error) {
	database, err := cmd.Flags().GetString("database")
	if err != nil {
		return nil, fmt.Errorf("[getFlags] error al obtener el nombre de la base de datos: %w", err)
	}

	return &Flags{
		databaseName: database,
	}, nil
}

func Run(cmd *cobra.Command, args []string) {
	flags, err := getFlags(cmd)
	if err != nil {
		cmd.Printf("Error al obtener las banderas: %v\n", err)
		return
	}

	cmd.Println("Flags obtenidos:")
	cmd.Println("Nombre de la base de datos:", flags.databaseName)

	for _, arg := range args {
		cmd.Println("Argumento:", arg)
	}

	cmd.Println("Comando de backup ejecutado")
}
