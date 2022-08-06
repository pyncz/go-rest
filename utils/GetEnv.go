package utils

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("You must set your '" + key + "' environmental variable")
	}
	return value
}
