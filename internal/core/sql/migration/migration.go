package migration

import (
	"database/sql"
	"fmt"
)

func Migrate(db *sql.DB) error {
	if err := CreateTables(db); err != nil {
		return fmt.Errorf("[Migrate] Error creating tables: %w", err)
	}

	return nil
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		level TEXT NOT NULL,
		message TEXT NOT NULL,
		context TEXT,
		exception TEXT
	);`)

	if err != nil {
		return fmt.Errorf("[CreateTables] Error creating logs table: %w", err)
	}

	return nil
}
