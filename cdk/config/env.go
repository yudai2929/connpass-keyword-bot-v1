package config

import (
	"os"

	"github.com/joho/godotenv"
)

func SetupEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}

	env := os.Getenv("ENVIRONMENT")

	if env != "dev" && env != "" {
		return nil
	}

	return err
}
