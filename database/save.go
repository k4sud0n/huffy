package database

import (
	"database/sql"
	"fmt"
	"time"
)

func SaveData(db *sql.DB, title string, date time.Time) error {
	var existingCount int
	err := db.QueryRow("SELECT COUNT(*) FROM ARTICLE WHERE TITLE = ? AND DATE = ?", title, date).Scan(&existingCount)
	if err != nil {
		return fmt.Errorf("error checking duplicates: %w", err)
	}

	if existingCount > 0 {
		return nil
	}

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
