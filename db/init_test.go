package db_test

import (
	"os"
	"testing"

	"github.com/jmoeser/go-git-sync/db"
)

func TestDatabaseInit(t *testing.T) {
	db.DatabaseInit("ggs.db")
	_, err := os.Stat("./ggs.db")
	if os.IsNotExist(err) {
		t.Error("Database has not been created")
	}
	defer os.Remove("./ggs.db")
}
