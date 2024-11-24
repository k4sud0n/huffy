package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	query1 := `
	CREATE TABLE IF NOT EXISTS NOTICE (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		TITLE TEXT NOT NULL,
		LINK TEXT NOT NULL,
		"DATE" DATETIME NOT NULL
		);
	`

	query2 := `
	CREATE TABLE IF NOT EXISTS MENU (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		CONTENT TEXT NOT NULL
	);
	`

	_, err = db.Exec(query1)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(query2)
	if err != nil {
		return nil, err
	}

	return db, nil
}
