package sqlite

import (
	"database/sql"
	"fmt"

	"criteria.mx/scripts/internal/core/sql/migration"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitSqlite() error {
	DB, err := sql.Open("sqlite3", "./toolkit.db")
	if err != nil {
		return fmt.Errorf("[Init] Error opening SQLite database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("[Init] Error pinging SQLite database: %w", err)
	}

	if err := migration.Migrate(DB); err != nil {
		return fmt.Errorf("[Init] Error during migration: %w", err)
	}

	return nil
}
