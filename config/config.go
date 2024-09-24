package config

import (
	"os"
)

// Config holds the configuration structure
type Config struct {
	DB       DatabaseConfig `env:"DB"`
	Firebase FirebaseConfig `env:"FIREBASE"`
}

// DatabaseConfig holds database-related configurations
type DatabaseConfig struct {
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
	Password string `env:"DB_PASSWORD"`
	User     string `env:"DB_USER"`
}

// FirebaseConfig holds Firebase-related configurations
type FirebaseConfig struct {
	BaseURL string `env:"FIREBASE_BASEURL"`
	APIKey  string `env:"FIREBASE_APIKEY"`
}

// LoadConfig reads the configuration from environment variables
func LoadConfig() (*Config, error) {
	cfg := &Config{
		DB: DatabaseConfig{
			Host:     getEnv("DB_HOST", ""),
			Name:     getEnv("DB_NAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
			User:     getEnv("DB_USER", ""),
		},
		Firebase: FirebaseConfig{
			BaseURL: getEnv("FIREBASE_BASEURL", ""),
			APIKey:  getEnv("FIREBASE_APIKEY", ""),
		},
	}

	return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
