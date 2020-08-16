package db_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/jmoeser/go-git-sync/db"
)

func TestUpdateSyncStatus(t *testing.T) {
	dbPath := "ggs.db"
	hash := "5ee116a"
	repo := "https://github.com/jmoeser/go-git-sync.git"

	db.DatabaseInit(dbPath)
	defer os.Remove(dbPath)

	db.UpdateSyncStatus(dbPath, repo, hash, "http://consul")

	type SyncStatus struct {
		Repo        string
		Hash        string
		Destination string
		// Date
	}

	sqliteDB, _ := sql.Open("sqlite3", dbPath)
	defer sqliteDB.Close()

	var SyncStatuses []*SyncStatus

	statement, err := sqliteDB.Prepare("SELECT repo, hash, destination FROM syncStatus where hash = ?;")
	if err != nil {
		t.Error(err)
	}

	rows, err := statement.Query(hash)
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		s := &SyncStatus{}
		if err := rows.Scan(&s.Repo, &s.Hash, &s.Destination); err != nil {
			t.Error(err)
		}
		SyncStatuses = append(SyncStatuses, s)
	}

	syncStatus := SyncStatuses[0]
	if syncStatus.Repo != repo {
		t.Error("Repo is not correct!")
	}
	if syncStatus.Hash != hash {
		t.Error("Hash is not correct!")
	}
}
