package main

import (
	"log"

	"criteria.mx/scripts/internal/core"
	"criteria.mx/scripts/pkg"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️  No se encontró archivo .env, usando variables de entorno del sistema")
	}

	if err := pkg.InitLogger(); err != nil {
		log.Fatalf("⚠️  Error initializing logger: %v", err)
	}

	core.Run()
}
