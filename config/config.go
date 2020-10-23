package config

import (
	"os"
)

// AppConfig is app config
type AppConfig struct {
	App  string
	Port int
	Env  string
}

// GetConfigByEnv get app config by os env: "ENV"
func GetConfigByEnv() *AppConfig {
	env := os.Getenv("ENV")

	switch env {
	case "production":
		return productConfig
	case "stage":
		return stageConfig
	default:
		os.Setenv("ENV", "development")
		return devConfig
	}
}
