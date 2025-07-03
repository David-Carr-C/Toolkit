package main

import (
	"log"
	"runtime/debug"

	"criteria.mx/scripts/internal/core"
	"criteria.mx/scripts/internal/core/sql/sqlite"
	"criteria.mx/scripts/pkg"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			log.Fatalf("⚠️  Application panicked: %v\nStack trace:\n%s", r, stack)
		}
	}()

	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️  No se encontró archivo .env, usando variables de entorno del sistema")
	}

	if err := pkg.InitLogger(); err != nil {
		log.Fatalf("⚠️  Error initializing logger: %v", err)
	}

	if err := sqlite.InitSqlite(); err != nil {
		log.Fatalf("⚠️  Error initializing SQLite: %v", err)
	}

	core.Run()
}
