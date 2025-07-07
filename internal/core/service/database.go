package service

import (
	"criteria.mx/scripts/internal/core/config/db"
	"criteria.mx/scripts/internal/core/config/sv"
	"criteria.mx/scripts/internal/sql/repo"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

func reportIt(elements []string, errToSave error) {
	repository, err := repo.NewRepository(nil)
	if err != nil {
		panic(err)
	}

	_, err = repository.InsertLog(
		fmt.Sprintf("La base de datos [%s] no pudo ser respaldada debido al error [%s]", elements, errToSave.Error()),
		fmt.Sprintf("Error: %s", errToSave.Error()),
		repo.CRITICAL,
	)

	if err != nil {
		panic(err)
	}
}

func backupIt(
	wg *sync.WaitGroup,
	databaseToBackup,
	pathDestination string,
	config *db.DatabaseConfig,
	errors chan error,
) {
	defer wg.Done()

	if err := os.MkdirAll(pathDestination, os.ModePerm); err != nil {
		errors <- fmt.Errorf("[BackupIt] No se pudo crear el directorio: %w, el %s backup fue cancelado", err, databaseToBackup)
	}

	timeStamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("%s_%s.sql", databaseToBackup, timeStamp)
	outputPath := filepath.Join(pathDestination, fileName)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		errors <- fmt.Errorf("[BackupIt] No se pudo crear el archivo de respaldo: %w, el %s backup fue cancelado", err, databaseToBackup)
		return
	}

	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			errors <- fmt.Errorf("[BackupIt] No se pudo cerrar el archivo de respaldo: %w, el %s backup fue cancelado", err, databaseToBackup)
			return
		}
	}(outputFile)

	cmd := exec.Command(
		"mysqldump",
		"--single-transaction",
		"--skip-lock-tables",
		fmt.Sprintf("--user=%s", config.User),
		fmt.Sprintf("--host=%s", config.Host),
		fmt.Sprintf("--port=%d", config.Port),
		databaseToBackup,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", config.Password))
	cmd.Stdout = outputFile
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		errors <- fmt.Errorf("[BackupIt] Error al ejecutar mysqldump: %w, el %s backup fue cancelado", err, databaseToBackup)
		return
	}
}

func backupEvery(config *db.DatabaseConfig) error {
	_, path, err := sv.GetServerPath("destination")
	if err != nil {
		reportIt(config.BackupDbs, err)
		return err
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(config.BackupDbs))

	for _, element := range config.BackupDbs {
		wg.Add(1)
		go backupIt(&wg, element, path, config, errors)
	}

	wg.Wait()
	close(errors)

	var errs []error
	if len(errors) > 0 {
		for err := range errors {
			reportIt(config.BackupDbs, err)
			errs = append(errs, err)
		}

		return fmt.Errorf("[BackupEvery] El backup fallo debido a: %s", errs)
	}

	return nil
}

func Database(database string) error {
	config, err := db.GetDatabaseStruct(database)
	if err != nil {
		return fmt.Errorf("[Database] Error al obtener la configuración de la base de datos: %s", err)
	}

	err = backupEvery(config)
	if err != nil {
		return fmt.Errorf("[Database] Error al realizar el respaldo %s", err)
	}

	return nil
}
