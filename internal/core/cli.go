package core

import (
	"criteria.mx/scripts/internal/backup"
	"criteria.mx/scripts/internal/database"
	"criteria.mx/scripts/internal/toolkit"
)

func Run() {
	cmd := toolkit.NewToolkit()
	cmd.AddCommand(backup.NewBackup())
	cmd.AddCommand(database.NewDatabase())
	cmd.Execute()
}
