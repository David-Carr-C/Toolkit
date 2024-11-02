package models

import (
	"github.com/spf13/cobra"
)

type CommandLineInterface struct {
	rootCmd    *cobra.Command
	compareCmd *cobra.Command
	updateCmd  *cobra.Command
	rsyncCmd   *cobra.Command
}

func Constructor() *CommandLineInterface {
	cli := CommandLineInterface{}
	cli.init()
	cli.rootCmd.Execute()
	return &cli
}

func (cli *CommandLineInterface) init() {
	cli.rootCmd = &cobra.Command{
		Use:     "toolkit",
		Aliases: []string{"tk"},
		Short:   "Aplicación de terminal para la gestión de servidores y su automatización",
		Long:    "Automatiza tareas diarias de administración de servidores, como la creación de backups, la comparación, la actualización y el versionamiento de bases de datos, entre otras.",
		Run:     cli.runRootCmd,
	}

	// cli.rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "Config file")

	cli.rsyncCmd = &cobra.Command{
		Use:   "backup",
		Short: "",
		Long:  "",
		Run:   nil,
	}

	cli.rsyncCmd.Flags().StringP("application", "a", "", "Application name")
	cli.rsyncCmd.Flags().StringP("destination", "d", "", "Destination path")

	cli.compareCmd = &cobra.Command{
		Use:   "compare",
		Short: "Compare two databases",
		Long:  "Compare two databases and generate a changelog.xml file",
		Run:   cli.runCompareCmd,
	}

	cli.updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update the changelog.xml file",
		Run:   cli.runUpdateCmd,
	}

	cli.rootCmd.AddCommand(cli.compareCmd)
	cli.rootCmd.AddCommand(cli.updateCmd)
}

/*
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
*/
