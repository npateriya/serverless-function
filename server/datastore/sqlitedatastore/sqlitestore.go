package sqlitedatastore

import (
	"database/sql"
	//	"fmt"
	"log"
	//	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	CDB *sql.DB
}

func sqliteStoreConnect(connectString string) *sqliteStore {
	db, err := sql.Open("sqlite3", "./serverless.db")
	if err != nil {
		log.Fatal(err)
	}
	return &sqliteStore{db}
}
