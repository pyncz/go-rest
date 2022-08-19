package setup

import (
	"log"
	"os"
)

func Logger() *log.Logger {
	return log.New(os.Stdout, "[go-rest] ", log.LstdFlags)
}
