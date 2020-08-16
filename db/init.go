package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func DatabaseInit(databasePath string) {
	log.Debug().Msg("Creating sqlite database")
	file, err := os.Create(databasePath)
	if err != nil {
		log.Fatal().Err(err)
	}
	file.Close()

	sqliteDatabase, _ := sql.Open("sqlite3", databasePath)
	defer sqliteDatabase.Close()
	createTables(sqliteDatabase)
}

func createTables(db *sql.DB) {

	syncStatusTableSQL := `CREATE TABLE syncStatus (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"repo" TEXT,
		"hash" TEXT,
		"destination" TEXT,
		"syncedAt" DATETIME DEFAULT CURRENT_TIMESTAMP
	  );`

	log.Debug().Msg("Creating syncStatus database")
	statement, err := db.Prepare(syncStatusTableSQL)
	if err != nil {
		log.Fatal().Err(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Debug().Msg("Created syncStatus database")
}
