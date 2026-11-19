package constants

import (
	"log"
	"os"
)

const (
	USER_INDEX = "user"
	POST_INDEX = "post"
)

// Connection settings and secrets are read from environment variables,
// so they are never stored in the source code.
var (
	ES_URL      = mustGetenv("ES_URL")
	ES_USERNAME = mustGetenv("ES_USERNAME")
	ES_PASSWORD = mustGetenv("ES_PASSWORD")
	GCS_BUCKET  = mustGetenv("GCS_BUCKET")
	JWT_SECRET  = mustGetenv("JWT_SECRET")
)

func mustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing required environment variable %s", key)
	}
	return value
}
