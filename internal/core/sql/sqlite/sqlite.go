package sqlite

import (
	"database/sql"
	"fmt"

	"criteria.mx/scripts/internal/core/sql/migration"
	"criteria.mx/scripts/internal/core/sql/repo"
	_ "github.com/mattn/go-sqlite3"
)

func InitSqlite() error {
	db, err := sql.Open("sqlite3", "./toolkit.db")
	if err != nil {
		return fmt.Errorf("[InitSqlite] Error opening SQLite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("[InitSqlite] Error pinging SQLite database: %w", err)
	}

	if err := migration.NewMigration(db); err != nil {
		return fmt.Errorf("[InitSqlite] Error during migration: %w", err)
	}

	if _, err := repo.NewRepository(db); err != nil {
		return fmt.Errorf("[InitSqlite] Error initializing repository: %w", err)
	}

	return nil
}
