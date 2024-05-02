package storage

import (
	"database/sql"
	"fitness-api/cmd/utils"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	utils.LoadEnv() // Load environment variables
	var env = utils.GetEnvVars()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_NAME))

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	log.Println("Successfully connected to the database!")
}

func GetDB() *sql.DB {
	return db
}
