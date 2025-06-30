package backup

import "github.com/spf13/cobra"

func Run(cmd *cobra.Command, args []string) {
	cmd.Println("Comando de backup ejecutado")
	panic("Función de backup no implementada")
}
