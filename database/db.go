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

	createNoticeTableQuery := `
	CREATE TABLE IF NOT EXISTS NOTICE (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		TITLE TEXT NOT NULL,
		LINK TEXT NOT NULL,
		"DATE" DATETIME NOT NULL
	);
	`

	createMenuTableQuery := `
	CREATE TABLE IF NOT EXISTS MENU (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		"DATE" TEXT NOT NULL,
		CONTENT TEXT NOT NULL
	);
	`

	_, err = db.Exec(createNoticeTableQuery)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createMenuTableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}
