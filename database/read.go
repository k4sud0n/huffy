package database

import (
	"database/sql"
	"fmt"
)

type Notice struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
	Date  string `json:"date"`
}

func ReadData(db *sql.DB) ([]Notice, error) {
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
