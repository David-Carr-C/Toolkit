package service

import (
	"criteria.mx/scripts/internal/core/config/db"
	"criteria.mx/scripts/internal/core/config/sv"
	"criteria.mx/scripts/internal/core/service/log"
	"criteria.mx/scripts/internal/sql/repo"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

type DatabaseService struct {
	sourceHost      string
	sourcePath      string
	destinationHost string
	destinationPath string
	isTunnelNeeded  bool
	isLocalHost     bool
	tunnelCmd       *exec.Cmd
	databaseConfig  *db.DatabaseConfig
	wg              *sync.WaitGroup
}

type backupResult struct {
	element string
	err     error
}

func NewDatabaseService(database, source, destination string) (*DatabaseService, error) {
	if database == "" {
		return nil, service.ReportIt(
			database,
			fmt.Errorf("[NewDatabaseService] database is required"),
			repo.ERROR,
		)
	}

	if source == "" {
		return nil, service.ReportIt(
			source,
			fmt.Errorf("[NewDatabaseService] source is required"),
			repo.ERROR,
		)
	}

	if destination == "" {
		return nil, service.ReportIt(
			destination,
			fmt.Errorf("[NewDatabaseService] destination is required"),
			repo.ERROR,
		)
	}

	config, err := db.GetDatabaseStruct(database)
	if err != nil {
		return nil, service.ReportIt(
			source,
			fmt.Errorf("[NewDatabaseService] Error al obtener la configuración de la base de datos: %s", err),
			repo.ERROR,
		)
	}

	destinationHost, destinationPath, err := sv.GetServerPath(destination)
	if err != nil {
		return nil, service.ReportIt(
			source,
			fmt.Errorf("[NewDatabaseService] Error al obtener la base de datos: %s", err),
			repo.ERROR,
		)
	}

	sourceHost, sourcePath, err := sv.GetServerPath(source)
	if err != nil {
		return nil, service.ReportIt(
			source,
			fmt.Errorf("[NewDatabaseService] Error al obtener la base de datos: %s", err),
			repo.ERROR,
		)
	}

	result := new(DatabaseService)
	result.destinationHost = destinationHost
	result.destinationPath = destinationPath
	result.sourceHost = sourceHost
	result.sourcePath = sourcePath
	result.databaseConfig = config
	result.isTunnelNeeded = config.Tunnel
	result.wg = &sync.WaitGroup{}

	if destinationHost == "localhost" {
		result.isLocalHost = true
	}

	return result, nil
}

func (d *DatabaseService) databaseSuccess(databaseName string) {
	// TODO: Implementar la lógica para manejar el éxito del backup de la base de datos
}

func (d *DatabaseService) stopTunnel() error {
	if d.tunnelCmd != nil && d.tunnelCmd.Process != nil {
		if err := d.tunnelCmd.Process.Kill(); err != nil {
			return fmt.Errorf("[stopTunnel] Error al detener el túnel: %w", err)
		}
	}

	return nil
}

func (d *DatabaseService) crateTunnel() error {
	destinationRandPort := 3307 + rand.Intn(100)
	cmd := exec.Command(
		"ssh",
		"-N",
		"-L",
		fmt.Sprintf("%d:%s:%d", destinationRandPort, d.destinationHost, d.databaseConfig.Port),
		d.sourceHost,
	)

	d.databaseConfig.Port = destinationRandPort
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		d.tunnelCmd = cmd
		if err_ := d.stopTunnel(); err_ != nil {
			return fmt.Errorf("[crateTunnel] Error al iniciar túnel: %w y al detener el túnel:  %w", err, err_)
		}
		return fmt.Errorf("[crateTunnel] Error al iniciar el túnel: %w", err)
	}

	d.tunnelCmd = cmd
	time.Sleep(2 * time.Second)

	// TODO: Agregar un mecanismo de verificación para asegurar que el túnel se ha establecido correctamente

	return nil
}

func (d *DatabaseService) backupIt(databaseToBackup string, errors chan backupResult) {
	defer d.wg.Done()

	if err := os.MkdirAll(d.destinationPath, os.ModePerm); err != nil {
		errors <- backupResult{
			err:     fmt.Errorf("[BackupIt] No se pudo crear el directorio: %w, el %s backup fue cancelado", err, databaseToBackup),
			element: databaseToBackup,
		}
		return
	}

	timeStamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("%s_%s.sql", databaseToBackup, timeStamp)
	outputPath := filepath.Join(d.destinationPath, fileName)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		errors <- backupResult{
			err:     fmt.Errorf("[BackupIt] No se pudo crear el archivo de respaldo: %w, el %s backup fue cancelado", err, databaseToBackup),
			element: databaseToBackup,
		}
		return
	}

	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			errors <- backupResult{
				err:     fmt.Errorf("[BackupIt] No se pudo cerrar el archivo de respaldo: %w, el %s backup fue cancelado", err, databaseToBackup),
				element: databaseToBackup,
			}
			return
		}
	}(outputFile)

	cmd := exec.Command(
		"mysqldump",
		"--single-transaction", // Evita bloqueos de tablas durante la copia
		"--quick",              // Descarga por lotes (evita memoria)
		"--compress",           // Compresión cliente/servidor
		"--add-drop-database",  // Añade DROP DATABASE antes de CREATE
		"--add-drop-table",     // Añade DROP TABLE antes de CREATE TABLE
		"--create-options",     // Incluye opciones CREATE TABLE completas
		"--extended-insert",    // Inserciones multi-fila (más compacto)
		"--routines",           // Incluye stored procedures
		"--triggers",           // Incluye triggers
		"--events",             // Incluye eventos programados
		"--skip-comments",      // Excluye comentarios innecesarios
		"--skip-add-locks",     // Mejora rendimiento
		fmt.Sprintf("--user=%s", d.databaseConfig.User),
		fmt.Sprintf("--host=%s", d.databaseConfig.Host),
		fmt.Sprintf("--port=%d", d.databaseConfig.Port),
		databaseToBackup,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", d.databaseConfig.Password))
	cmd.Stdout = outputFile
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		errors <- backupResult{
			err:     fmt.Errorf("[BackupIt] Error al ejecutar mysqldump: %w, el %s backup fue cancelado", err, databaseToBackup),
			element: databaseToBackup,
		}

		err = os.Remove(outputPath)
		if err != nil {
			errors <- backupResult{
				err:     fmt.Errorf("[BackupIt] No se pudo eliminar el archivo malformado de respaldo: %w, el %s backup fue cancelado", err, databaseToBackup),
				element: databaseToBackup,
			}
			return
		}

		return
	}
}

func (d *DatabaseService) Exec() error {
	errors := make(chan backupResult, len(d.databaseConfig.BackupDbs))

	if d.isTunnelNeeded {
		err := d.crateTunnel()
		if err != nil {
			for _, element := range d.databaseConfig.BackupDbs {
				_ = service.ReportIt(
					element,
					fmt.Errorf("[BackupIt] Error al crear el túnel: %w, el %s backup fue cancelado", err, element),
					repo.ERROR,
				)
			}

			return fmt.Errorf("[BackupIt] Error al crear el túnel: %w, el %s backup fue cancelado", err, d.databaseConfig.BackupDbs)
		}
	}

	for _, element := range d.databaseConfig.BackupDbs {
		d.wg.Add(1)
		go d.backupIt(element, errors)
	}

	d.wg.Wait()
	close(errors)
	var persistErrors []backupResult
	for err := range errors {
		persistErrors = append(persistErrors, err)
	}

	if d.isTunnelNeeded {
		if err := d.stopTunnel(); err != nil {
			_ = service.ReportIt(
				"tunnel",
				fmt.Errorf("[BackupIt] Error al detener el túnel: %w", err),
				repo.ERROR,
			)
		}
	}

	for _, dbName := range d.databaseConfig.BackupDbs {
		found := false
		for _, err := range persistErrors {
			if err.element == dbName {
				found = true
				break
			}
		}
		if found {
			continue
		}

		d.databaseSuccess(dbName)
	}

	var errs []error
	if len(persistErrors) > 0 {
		for _, err := range persistErrors {
			errs = append(errs, service.ReportIt(
				err.element,
				err.err,
				repo.CRITICAL,
			))
		}

		return fmt.Errorf("[BackupEvery] El backup fallo debido a: %s", errs)
	}

	return nil
}
