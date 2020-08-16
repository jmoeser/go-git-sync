package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func UpdateSyncStatus(databasePath string, repo string, hash string, destination string) {

	db, _ := sql.Open("sqlite3", databasePath)
	defer db.Close()

	insertSQL := `INSERT INTO syncStatus(repo, hash, destination) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal().Err(err)
	}
	_, err = statement.Exec(repo, hash, destination)
	if err != nil {
		log.Fatal().Err(err)
	}
}
