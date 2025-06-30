package pkg

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	logFile = os.Getenv("LOG_FILE")
)

func InitLogger() error {
	if logFile == "" {
		return fmt.Errorf("[InitLogger] LOG_FILE environment variable is not set")
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("[InitLogger] Error opening log file: %v", err)
	}

	log.SetOutput(file)
	return nil
}
