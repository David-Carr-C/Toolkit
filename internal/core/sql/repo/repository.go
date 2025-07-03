package repo

import (
	"database/sql"
	"fmt"

	"criteria.mx/scripts/internal/core/sql/model"
)

var repo *RepositoryLog

type RepositoryLog struct {
	sqlite *sql.DB
}

func NewRepository(db *sql.DB) (*RepositoryLog, error) {
	if db == nil && repo == nil {
		return nil, fmt.Errorf("[NewRepository] database connection is nil")
	}

	if repo == nil {
		repo = &RepositoryLog{sqlite: db}
	}

	return repo, nil
}

func (r *RepositoryLog) InsertLog(message, exception, level string) (int64, error) {
	result, err := r.sqlite.Exec(`INSERT INTO log (message, exception, level) VALUES (?, ?, ?)`, message, exception, level)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("[InsertLog] Error getting last insert ID: %w", err)
	}

	return lastID, nil
}

func (r *RepositoryLog) InsertSource(idLog int64, server string) (int64, error) {
	result, err := r.sqlite.Exec(`INSERT INTO source (id_log, server) VALUES (?, ?)`, idLog, server)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("[InsertSource] Error getting last insert ID: %w", err)
	}

	return lastID, nil
}

func (r *RepositoryLog) InsertDestination(idLog int64, server, path string) (int64, error) {
	result, err := r.sqlite.Exec(`INSERT INTO destination (id_log, server, path) VALUES (?, ?, ?)`, idLog, server, path)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("[InsertDestination] Error getting last insert ID: %w", err)
	}

	return lastID, nil
}

func (r *RepositoryLog) InsertBackup(idLog int64, sha256 string) (int64, error) {
	result, err := r.sqlite.Exec(`INSERT INTO backup (id_log, sha256) VALUES (?, ?)`, idLog, sha256)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("[InsertBackup] Error getting last insert ID: %w", err)
	}

	return lastID, nil
}

func (r *RepositoryLog) InsertDatabase(idLog int64, sha256 string) (int64, error) {
	result, err := r.sqlite.Exec(`INSERT INTO database (id_log, sha256) VALUES (?, ?)`, idLog, sha256)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("[InsertDatabase] Error getting last insert ID: %w", err)
	}

	return lastID, nil
}

func (r *RepositoryLog) UpdateSynchronized(idLog int64, synchronized bool) error {
	_, err := r.sqlite.Exec(`UPDATE log SET synchronized = ? WHERE id = ?`, synchronized, idLog)
	if err != nil {
		return fmt.Errorf("[UpdateSynchronized] Error updating synchronized status for log ID %d: %w", idLog, err)
	}

	return nil
}

// Retorna todos los logs o si se especifica un ID, retorna el log correspondiente.
// Parámetros:
// - idLog ID del log a buscar, si es 0 retorna todos los logs
// - limit Limite de logs a retornar, si es 0 retorna todos los logs
// Retorna:
// - *model.Log Retorna un puntero al modelo Log con los datos del log
func (r *RepositoryLog) GetLog(idLog, limit int64) (*model.Log, error) {
	query := `SELECT l.id, l.message, l.exception, l.level, l.created_at, l.synchronized,
				s.server AS source_server,
				d.server AS destination_server, d.path AS destination_path,
				b.sha256 AS backup_sha256,
				db.sha256 AS database_sha256
			FROM log l
				LEFT JOIN source s ON l.id = s.id_log
				LEFT JOIN destination d ON l.id = d.id_log
				LEFT JOIN backup b ON l.id = b.id_log
				LEFT JOIN database db ON l.id = db.id_log`

	if idLog > 0 {
		query += fmt.Sprintf(" WHERE l.id = %d", idLog)
	} else if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}

	query += ` ORDER BY l.id DESC`

	row := r.sqlite.QueryRow(query, idLog)
	var log model.Log
	err := row.Scan(
		&log.ID,
		&log.Message,
		&log.Exception,
		&log.Level,
		&log.CreatedAt,
		&log.SourceServer,
		&log.DestinationServer,
		&log.DestinationPath,
		&log.BackupSHA256,
		&log.DatabaseSHA256,
		&log.Synchronized,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("[GetLog] No log found with ID %d: %w", idLog, err)
		}
		return nil, fmt.Errorf("[GetLog] Error retrieving log with ID %d: %w", idLog, err)
	}

	return &log, nil
}
