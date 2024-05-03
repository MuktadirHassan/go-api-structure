package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type EnvVars struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

var envVars *EnvVars

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	// Validate required environment variables
	validateRequiredEnvVars(
		[]string{
			"DB_HOST",
			"DB_PORT",
			"DB_USER",
			"DB_PASSWORD",
			"DB_NAME",
		})

	envVars = &EnvVars{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

}

func GetEnvVars() *EnvVars {
	return envVars
}

func validateRequiredEnvVars(req []string) {
	var missingEnvVars []string
	for _, envVar := range req {
		if os.Getenv(envVar) == "" {
			missingEnvVars = append(missingEnvVars, envVar)
		} else {
			log.Debug().Msgf("Found environment variable: %s", envVar)
		}
	}

	if len(missingEnvVars) > 0 {
		log.Fatal().Msgf("Missing required environment variables: %v", missingEnvVars)
	}

}
