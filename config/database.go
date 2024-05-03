package config

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var DB *sql.DB

func InitializeDB() {
	env := GetEnvVars()

	DB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_NAME))

	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to the database")
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal().Err(err).Msg("Error pinging the database")
	}

	log.Info().Msg("Successfully connected to the database")
}
