package main

import (
	"criteria.mx/scripts/internal/core/config/db"
	"criteria.mx/scripts/internal/core/config/dir"
	"criteria.mx/scripts/internal/core/config/sv"
	"testing"
)

func TestDbConfig(t *testing.T) {
	databases, err := db.GetAllDatabases()
	if err != nil {
		t.Errorf("[testDbConfig] Error al obtener la configuracion de databases: %s", err.Error())
	}

	t.Log(databases)
}

func TestDirConfig(t *testing.T) {
	projects, err := dir.GetAllProjects()
	if err != nil {
		t.Errorf("[testDbConfig] Error al obtener la configuracion de databases: %s", err.Error())
	}

	t.Log(projects)
}

func TestSvConfig(t *testing.T) {
	servers, err := sv.GetAllServers()
	if err != nil {
		t.Errorf("[testDbConfig] Error al obtener la configuracion de databases: %s", err.Error())
	}

	t.Log(servers)
}

func BenchmarkDbConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := db.GetAllDatabases()
		if err != nil {
			b.Errorf("Error al obtener la configuracion de databases: %s", err.Error())
		}
	}
}

func BenchmarkDirConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := db.GetAllDatabases()
		if err != nil {
			b.Errorf("Error al obtener la configuracion de databases: %s", err.Error())
		}
	}
}

func BenchmarkSvConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := db.GetAllDatabases()
		if err != nil {
			b.Errorf("Error al obtener la configuracion de databases: %s", err.Error())
		}
	}
}
