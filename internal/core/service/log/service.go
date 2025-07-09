package service

import (
	"criteria.mx/scripts/internal/sql/repo"
	"fmt"
)

func ReportIt(element string, errToSave error, logLevel string) error {
	repository, err := repo.NewRepository(nil)
	if err != nil {
		fmt.Printf("[reportIt] Ocurrio un error intentando reportar a la base de datos: %v\n", err)
		panic(err)
	}

	switch logLevel {
	case repo.INFO:
	case repo.DEBUG:
	case repo.WARNING:
	case repo.ERROR:
	case repo.CRITICAL:
	default:
		logLevel = repo.ERROR
	}

	_, err = repository.InsertLog(
		element,
		fmt.Sprintf("No pudo ser respaldado el elemento [%s] debido al error [%s]", element, errToSave.Error()),
		fmt.Sprintf("Error: %s", errToSave.Error()),
		logLevel,
	)

	if err != nil {
		fmt.Printf("[reportIt] Ocurrio un error intentando reportar a la base de datos: %v\n", err)
		panic(err)
	}

	return errToSave
}
