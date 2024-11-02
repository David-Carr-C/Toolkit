package models

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Ocupar Ansible, systemctl, liquibase, awx/other

func init() {
	rootCmd := &cobra.Command{
		Use:     "toolkit",
		Aliases: []string{"tk"},
		Short:   "Aplicación de terminal para la gestión de servidores y su automatización",
		Long:    "Automatiza tareas diarias de administración de servidores, como la creación de backups, la comparación, la actualización y el versionamiento de bases de datos, entre otras.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from cobra")

		},
	}

	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "Respaldos de archivos y bases de datos",
		Long:  "Crea respaldos de archivos y bases de datos dependiendo de que aplicación se elija",
		Run: func(cmd *cobra.Command, args []string) {
			application, _ := cmd.Flags().GetString("application")
			destination, _ := cmd.Flags().GetString("destination")

			switch application {
			case "wordpress":
				fmt.Println("Backuping wordpress")
			case "magento":
				fmt.Println("Backuping magento")
			}

			fmt.Println("Backuping server to ", destination)
		},
	}

	backupCmd.Flags().StringP("application", "a", "", "Application name")
	backupCmd.Flags().StringP("destination", "d", "", "Destination path")

	rootCmd.AddCommand(backupCmd)
}

func (cli *CommandLineInterface) runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Hello from cobra")

}

func (cli *CommandLineInterface) runCompareCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Comparing databases asdasdasdasdsa")
	// database := args[0]
	// database2 := args[1]
	// fmt.Println("Comparing databases, ", database, database2)
	// // mysql dump
	// o := exec.Command("mysqldump", "-u", "root", "-p", database)
	// output, _ := o.Output()
	// fmt.Println(string(output))
}

func (cli *CommandLineInterface) dumpDatabasesCmd(cmd *cobra.Command, args []string) {
	// database1 := args[0]
	// database2 := args[1]
	// fmt.Println("Dumping databases", database1, database2)

	// cmd1 := exec.Command("mysqldump", "-u", "root", "-p", database1)
	// output, _ := cmd1.Output()
	// fmt.Println(string(output))
}

func (cli *CommandLineInterface) runUpdateCmd(cmd *cobra.Command, args []string) {
	// changeLogPath := "changelog.xml"
	// cmd1 := exec.Command("liquibase", "--changeLogFile="+changeLogPath, "diffChangeLog")
	// output, err := cmd1.Output()
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// fmt.Println("Done")
	// fmt.Println(string(output))
}
