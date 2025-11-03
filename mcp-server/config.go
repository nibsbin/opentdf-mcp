package main

import (
	"os"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file if present. Ignore errors because env vars may be
	// provided via the environment in production.
	_ = godotenv.Load()
}

func getPlatformEndpoint() string {
	if endpoint := os.Getenv("OPENTDF_PLATFORM_ENDPOINT"); endpoint != "" {
		return endpoint
	}
	return "http://localhost:8080"
}

func getClientID() string {
	if clientID := os.Getenv("OPENTDF_CLIENT_ID"); clientID != "" {
		return clientID
	}
	return "opentdf-sdk"
}

func getClientSecret() string {
	if secret := os.Getenv("OPENTDF_CLIENT_SECRET"); secret != "" {
		return secret
	}
	return "secret"
}
