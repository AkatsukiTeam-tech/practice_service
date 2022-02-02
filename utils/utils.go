package utils

import (
	"log"
	"os"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	log.Println("value: ", value)
	if value != "" {
		return value
	}
	return defaultValue
}
