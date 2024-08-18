package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"

	"gamch1k.org/ggit/utils"
)


func connect() (*sql.DB, error) {
	rootDir := utils.GetEnvVariable("ROOT_PATH")
	db, err := sql.Open("sqlite", rootDir + "/database/database.db")
	if utils.ErrorHandler(err) {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}


func CreateFields() {
	db, err := connect()
	if utils.ErrorHandler(err) {
		return
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS commits (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hash TEXT NOT NULL UNIQUE,
		message TEXT NOT NULL,
		created_at DATETIME NOT NULL
	);
	`

	_, err = db.Exec(sqlStmt)
	if utils.ErrorHandler(err) {
		return
	}
}
