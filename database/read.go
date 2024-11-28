package database

import (
	"database/sql"
	"fmt"
)

func ReadMenu(db *sql.DB) ([]Menu, error) {
	query := `SELECT ID, "DATE", CONTENT FROM MENU WHERE substr("DATE", 1, 5) = strftime('%m/%d', 'now')`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var menus []Menu
	for rows.Next() {
		var menu Menu
		err := rows.Scan(&menu.ID, &menu.Date, &menu.Content)
		if err != nil {
			return nil, fmt.Errorf("error scanning notice row: %w", err)
		}
		menus = append(menus, menu)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return menus, nil
}

func ReadNotice(db *sql.DB) ([]Notice, error) {
	query := `SELECT ID, TITLE, LINK, strftime('%Y-%m-%d', "DATE") AS "DATE" FROM NOTICE`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notices []Notice
	for rows.Next() {
		var notice Notice
		err := rows.Scan(&notice.ID, &notice.Title, &notice.Link, &notice.Date)
		if err != nil {
			return nil, fmt.Errorf("error scanning notice row: %w", err)
		}
		notices = append(notices, notice)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return notices, nil
}
