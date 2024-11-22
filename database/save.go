package database

import (
	"database/sql"
	"time"
)

func SaveData(db *sql.DB, title string, date time.Time) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO ARTICLE (TITLE, "DATE") VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, date)
	if err != nil {
		return err
	}

	return tx.Commit()
}
