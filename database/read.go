package database

import (
	"database/sql"
	"fmt"
)

func ReadMenu(db *sql.DB, parameter string) ([]Menu, error) {
	query := `SELECT ID, "DATE", CONTENT
			FROM MENU
			WHERE "DATE" = strftime('%Y/%m/%d', datetime('now', '+09:00'))
  			AND LOCATION = ?;`

	rows, err := db.Query(query, parameter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
	query := `SELECT ID, TITLE, LINK, strftime('%Y-%m-%d', datetime("DATE", '+09:00')) AS "DATE" 
          FROM NOTICE
          ORDER BY "DATE" DESC
          LIMIT 5;`

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
