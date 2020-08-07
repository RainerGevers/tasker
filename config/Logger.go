package config

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	logger := log.New(os.Stdout, "Tasker ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds|log.LUTC)

	return logger
}
