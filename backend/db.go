package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var db *sql.DB

func initTurso() {
	url := os.Getenv("TURSO_URL")
	token := os.Getenv("TURSO_TOKEN")

	connectionString := fmt.Sprintf("%s?authToken=%s", url, token)
	_db, err := sql.Open("libsql", connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	db = _db
}
