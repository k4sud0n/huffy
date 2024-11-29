package database

import (
	"database/sql"
	"fmt"
	"time"
)

func SaveMenu(db *sql.DB, date string, location string, content string) error {
	var existingCount int
	err := db.QueryRow("SELECT COUNT(*) FROM MENU WHERE \"DATE\" = ? AND LOCATION = ?", date, location).Scan(&existingCount)
	if err != nil {
		return fmt.Errorf("error checking for duplicates: %w", err)
	}

	if existingCount > 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO MENU ("DATE", LOCATION, CONTENT) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(date, location, content)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func SaveNotice(db *sql.DB, title string, link string, date time.Time) error {
	var existingCount int
	err := db.QueryRow("SELECT COUNT(*) FROM NOTICE WHERE TITLE = ? AND DATE = ?", title, date).Scan(&existingCount)
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

	stmt, err := tx.Prepare(`INSERT INTO NOTICE (TITLE, LINK, "DATE") VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, link, date)
	if err != nil {
		return err
	}

	return tx.Commit()
}
