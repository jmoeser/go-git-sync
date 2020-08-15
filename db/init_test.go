package db

import (
	"os"
	"testing"
)

func TestDatabaseInit(t *testing.T) {
	DatabaseInit("ggs.db")
	_, err := os.Stat("./ggs.db")
	if os.IsNotExist(err) {
		t.Error("Database has not been created")
	}
	defer os.Remove("./ggs.db")
}
