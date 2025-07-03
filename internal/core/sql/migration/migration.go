package migration

import (
	"database/sql"
	"fmt"
)

type Migration struct {
	db *sql.DB
}

func NewMigration(db *sql.DB) error {
	migration := &Migration{db}

	err := migration.migrate()
	if err != nil {
		return fmt.Errorf("[NewMigration] Error during migration: %w", err)
	}

	return nil
}

func (s *Migration) migrate() error {
	if err := s.createLogTable(); err != nil {
		return fmt.Errorf("[Migrate] Error creating tables: %w", err)
	}

	if err := s.createSourceTable(); err != nil {
		return fmt.Errorf("[Migrate] Error creating remote table: %w", err)
	}

	if err := s.createDestinationTable(); err != nil {
		return fmt.Errorf("[Migrate] Error creating local table: %w", err)
	}

	if err := s.createBackupTable(); err != nil {
		return fmt.Errorf("[Migrate] Error creating repository table: %w", err)
	}

	if err := s.createDatabaseTable(); err != nil {
		return fmt.Errorf("[Migrate] Error creating repository log table: %w", err)
	}

	return nil
}

func (s *Migration) createLogTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		message TEXT NOT NULL,
		exception TEXT,
		level TEXT NOT NULL CHECK(level IN ('DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL'))
	);`)

	if err != nil {
		return fmt.Errorf("[CreateLogTable] Error creating log table: %w", err)
	}

	return nil
}

func (s *Migration) createSourceTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS source (
		id_log INTEGER NOT NULL,
		server TEXT NOT NULL,
		FOREIGN KEY (id_log) REFERENCES log(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`)

	if err != nil {
		return fmt.Errorf("[CreateSourceTable] Error creating source table: %w", err)
	}

	return nil
}

func (s *Migration) createDestinationTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS destination (
		id_log INTEGER NOT NULL,
		server TEXT NOT NULL,
		path TEXT NOT NULL,
		FOREIGN KEY (id_log) REFERENCES log(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`)

	if err != nil {
		return fmt.Errorf("[CreateDestinationTable] Error creating destination table: %w", err)
	}

	return nil
}

func (s *Migration) createBackupTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS backup (
		id_log INTEGER NOT NULL,
		sha256 TEXT NOT NULL,
		FOREIGN KEY (id_log) REFERENCES log(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`)

	if err != nil {
		return fmt.Errorf("[CreateBackupTable] Error creating repository table: %w", err)
	}

	return nil
}

func (s *Migration) createDatabaseTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS database (
		id_log INTEGER NOT NULL,
		sha256 TEXT NOT NULL,
		FOREIGN KEY (id_log) REFERENCES log(id) ON DELETE CASCADE ON UPDATE CASCADE
	);`)

	if err != nil {
		return fmt.Errorf("[CreateDatabaseTable] Error creating repository log table: %w", err)
	}

	return nil
}
