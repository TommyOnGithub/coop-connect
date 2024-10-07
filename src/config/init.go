package config

import (
	"log"
	"os"
)

var APP_ROOT string

func init() {
	APP_ROOT = os.Getenv("APP_ROOT")
	if APP_ROOT == "" {
		log.Fatal("APP_ROOT environment variable is not set")
	}
}
