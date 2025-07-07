package test

import (
	"criteria.mx/scripts/internal/core/config/db"
	"testing"
)

func testDbConfig(t *testing.T) {
	_, err := db.GetAllDatabases()
	if err != nil {
		t.Errorf("[testDbConfig] Error al obtener la configuracion de databases: %s", err.Error())
	}
}

func testDirConfig() {

}

func testSvConfig() {

}
