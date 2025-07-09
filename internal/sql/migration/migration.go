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
	if err := s.pragma(); err != nil {
		return fmt.Errorf("[Migrate] Error setting PRAGMA: %w", err)
	}

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

func (s *Migration) pragma() error {
	_, err := s.db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error enabling foreign keys: %w", err)
	}
	_, err = s.db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting journal mode: %w", err)
	}
	_, err = s.db.Exec("PRAGMA synchronous = NORMAL;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting synchronous mode: %w", err)
	}
	_, err = s.db.Exec("PRAGMA cache_size = 10000;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting cache size: %w", err)
	}
	_, err = s.db.Exec("PRAGMA busy_timeout = 5000;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting busy timeout: %w", err)
	}
	_, err = s.db.Exec("PRAGMA temp_store = MEMORY;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting temp store: %w", err)
	}
	_, err = s.db.Exec("PRAGMA page_size = 4096;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting page size: %w", err)
	}
	_, err = s.db.Exec("PRAGMA auto_vacuum = FULL;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting auto vacuum: %w", err)
	}
	_, err = s.db.Exec("PRAGMA secure_delete = ON;")
	if err != nil {
		return fmt.Errorf("[Pragma] Error setting secure delete: %w", err)
	}

	return nil
}

func (s *Migration) createLogTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT (datetime('now', 'localtime')),
		element TEXT NOT NULL,
		message TEXT NOT NULL,
		exception TEXT,
		synchronized BOOLEAN NOT NULL DEFAULT 0,
		level TEXT NOT NULL CHECK(level IN ('SUCCESS', 'DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL'))
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
